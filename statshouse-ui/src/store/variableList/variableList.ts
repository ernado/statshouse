// Copyright 2023 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

import { Store, useStore } from '../statshouse';
import {
  getTagDescription,
  isTagEnabled,
  isValidVariableName,
  loadAllMeta,
  promQLMetric,
  replaceVariable,
} from '../../view/utils';
import { apiMetricTagValuesFetch, MetricTagValueInfo } from '../../api/metricTagValues';
import { GET_PARAMS, isTagKey, METRIC_VALUE_BACKEND_VERSION, QueryWhat, TAG_KEY, TagKey } from '../../api/enum';
import { globalSettings } from '../../common/settings';
import { filterParamsArr } from '../../view/api';
import { useErrorStore } from '../errors';
import { deepClone, isNotNil, toNumber } from '../../common/helpers';
import { MetricMetaTag } from '../../api/metric';
import { getEmptyVariableParams } from '../../common/getEmptyVariableParams';
import { PlotKey, toIndexTag, toKeyTag, toPlotKey, VariableParams, VariableParamsLink } from '../../url/queryParams';
import { createStore } from '../createStore';

export function getEmptyVariable(): VariableItem {
  return { list: [], updated: false, loaded: false, more: false, tagMeta: undefined, keyLastRequest: '' };
}

export type VariableItem = {
  list: MetricTagValueInfo[];
  updated: boolean;
  loaded: boolean;
  more: boolean;
  tagMeta?: MetricMetaTag;
  keyLastRequest: string;
};

export type VariableListStore = {
  variables: Record<string, VariableItem>;
  tags: Record<PlotKey, Record<TagKey, VariableItem>>;
};

export const useVariableListStore = createStore<VariableListStore>((setState, getState) => {
  useStore.subscribe((state, prevState) => {
    if (
      prevState.params.dashboard?.dashboard_id !== state.params.dashboard?.dashboard_id ||
      prevState.params.plots !== state.params.plots
    ) {
      if (
        prevState.params.plots.some(
          (plot, indexPlot) =>
            !state.params.plots[indexPlot] || plot.metricName !== state.params.plots[indexPlot]?.metricName
        )
      ) {
        clearTagsAll();
      }
    }
    if (prevState.params !== state.params) {
      updateVariables(state);
      updateTags(state);
    }
    if (prevState.metricsMeta !== state.metricsMeta) {
      const variableItems = getState().variables;
      state.params.variables.forEach((variable) => {
        if (!variableItems[variable.name].tagMeta) {
          variable.link.forEach(([plotKey, tagKey]) => {
            const indexPlot = toNumber(plotKey);
            const indexTag = toIndexTag(tagKey);
            if (indexPlot != null && indexTag != null) {
              const meta = state.metricsMeta[state.params.plots[indexPlot].metricName];
              setState((variableState) => {
                if (variableState.variables[variable.name]) {
                  variableState.variables[variable.name].tagMeta = meta?.tags?.[indexTag];
                }
              });
            }
          });
        }
      });
    }
  });
  return {
    variables: {},
    tags: {},
  };
}, 'VariableListStore');
export function updateTags(state: Store) {
  const plotKey = toPlotKey(state.params.tabNum);
  const updated: TagKey[] = [];
  if (plotKey != null) {
    const tags = useVariableListStore.getState().tags;
    if (tags[plotKey]) {
      Object.entries(tags[plotKey]).forEach(([indexTag, tagInfo]) => {
        if (tagInfo.updated && isTagKey(indexTag)) {
          updated.push(indexTag);
        }
      });
    }
    updated.forEach((indexTag) => {
      updateTag(plotKey, indexTag);
    });
  }
}
export async function updateTag(plotKey: PlotKey, tagKey: TagKey) {
  useVariableListStore.setState((state) => {
    state.tags[plotKey] ??= {} as Record<TagKey, VariableItem>;
    state.tags[plotKey][tagKey] ??= getEmptyVariable();
    if (state.tags[plotKey]?.[tagKey]) {
      state.tags[plotKey][tagKey].loaded = true;
    }
  });
  const listTag = await loadTagList(plotKey, tagKey);
  useVariableListStore.setState((state) => {
    if (state.tags[plotKey]?.[tagKey]) {
      state.tags[plotKey][tagKey].list = listTag?.values ?? [];
      state.tags[plotKey][tagKey].more = listTag?.more ?? false;
      state.tags[plotKey][tagKey].tagMeta = listTag?.tagMeta;
      state.tags[plotKey][tagKey].loaded = false;
      state.tags[plotKey][tagKey].keyLastRequest = listTag?.keyLastRequest ?? '';
    }
  });
}

export function setUpdatedTag(plotKey: PlotKey, tagKey: TagKey | undefined, toggle: boolean) {
  if (tagKey == null) {
    return;
  }
  useVariableListStore.setState((state) => {
    state.tags[plotKey] ??= {} as Record<TagKey, VariableItem>;
    state.tags[plotKey][tagKey] ??= getEmptyVariable();
    state.tags[plotKey][tagKey].updated = toggle;
  });
  if (toggle) {
    updateTag(plotKey, tagKey);
  }
}

export function clearTags(indexPlot: number) {
  useVariableListStore.setState((state) => {
    delete state.tags[indexPlot];
  });
}

export function clearTagsAll() {
  useVariableListStore.setState((state) => {
    state.tags = {};
  });
}

export function updateVariables(store: Store) {
  const update: VariableParams[] = [];
  useVariableListStore.setState((state) => {
    const variables: Record<string, VariableItem> = {};
    store.params.variables.forEach((variable) => {
      variables[variable.name] = state.variables[variable.name] ?? getEmptyVariable();
      if (variables[variable.name].updated) {
        update.push(variable);
      }
    });
    state.variables = variables;
  });
  update.forEach(updateVariable);
}

export async function updateVariable(variableParam: VariableParams) {
  useVariableListStore.setState((state) => {
    if (state.variables[variableParam.name]) {
      state.variables[variableParam.name].loaded = true;
    }
  });
  const lists = await loadValuableList(variableParam);
  lists.forEach((listTag) => {
    if (listTag) {
      useVariableListStore.setState((state) => {
        const { plotKey, tagKey } = listTag;
        state.tags[plotKey] ??= {} as Record<TagKey, VariableItem>;
        state.tags[plotKey][tagKey] ??= getEmptyVariable();
        if (state.tags[plotKey]?.[tagKey]) {
          state.tags[plotKey][tagKey].list = deepClone(listTag?.values ?? []);
          state.tags[plotKey][tagKey].more = listTag?.more ?? false;
          state.tags[plotKey][tagKey].tagMeta = deepClone(listTag?.tagMeta);
          state.tags[plotKey][tagKey].loaded = false;
          state.tags[plotKey][tagKey].keyLastRequest = listTag?.keyLastRequest ?? '';
        }
      });
    }
  });
  const more = lists.some((l) => l.more);
  const tagMeta = lists[0]?.tagMeta;
  const list = Object.values(
    lists
      .flatMap((l) => l.values)
      .reduce(
        (res, t) => {
          if (res[t.value]) {
            res[t.value].count += t.count;
          } else {
            res[t.value] = { ...t };
          }
          return res;
        },
        {} as Record<string, MetricTagValueInfo>
      )
  );
  useVariableListStore.setState((state) => {
    if (state.variables[variableParam.name]) {
      state.variables[variableParam.name].list = list;
      state.variables[variableParam.name].loaded = false;
      state.variables[variableParam.name].more = more;
      state.variables[variableParam.name].tagMeta = tagMeta;
    }
  });
}

export async function loadValuableList(variableParam: VariableParams) {
  const lists = await Promise.all(
    variableParam.link.map(async ([indexPlot, indexTag]) => await loadTagList(indexPlot, indexTag))
  );
  return lists.filter(isNotNil);
}

export async function loadTagList(plotKey: PlotKey, tagKey: TagKey, limit = 25000) {
  const indexPlot = toNumber(plotKey);
  const indexTag = toIndexTag(tagKey);
  const store = useStore.getState();
  if (
    indexPlot == null ||
    indexTag == null ||
    !store.params.plots[indexPlot] ||
    store.params.plots[indexPlot]?.metricName === promQLMetric
  ) {
    return undefined;
  }
  if (!tagKey) {
    return undefined;
  }
  const plot = replaceVariable(plotKey, store.params.plots[indexPlot], store.params.variables);
  const otherFilterIn = { ...plot.filterIn };
  delete otherFilterIn[tagKey];
  const otherFilterNotIn = { ...plot.filterNotIn };
  delete otherFilterNotIn[tagKey];
  const requestKey = `variable_${indexPlot}-${plot.metricName}`;
  await store.loadMetricsMeta(plot.metricName);
  const tagMeta = useStore.getState().metricsMeta[plot.metricName]?.tags?.[indexTag];
  const params = {
    [GET_PARAMS.metricName]: plot.metricName,
    [GET_PARAMS.metricTagID]: tagKey,
    [GET_PARAMS.version]:
      globalSettings.disabled_v1 || plot.useV2 ? METRIC_VALUE_BACKEND_VERSION.v2 : METRIC_VALUE_BACKEND_VERSION.v1,
    [GET_PARAMS.numResults]: limit.toString(),
    [GET_PARAMS.fromTime]: store.timeRange.from.toString(),
    [GET_PARAMS.toTime]: (store.timeRange.to + 1).toString(),
    [GET_PARAMS.metricFilter]: filterParamsArr(otherFilterIn, otherFilterNotIn),
    [GET_PARAMS.metricWhat]: plot.what.slice() as QueryWhat[],
  };
  const keyLastRequest = JSON.stringify(params);
  const lastTag = useVariableListStore.getState().tags[plotKey]?.[tagKey];
  if (lastTag && lastTag.keyLastRequest === keyLastRequest) {
    return {
      plotKey,
      tagKey,
      keyLastRequest: lastTag.keyLastRequest,
      values: lastTag.list,
      more: lastTag.more,
      tagMeta: lastTag.tagMeta,
    };
  }
  const { response, error } = await apiMetricTagValuesFetch(params, requestKey);
  if (response) {
    return {
      plotKey,
      tagKey,
      keyLastRequest,
      values: response.data.tag_values.slice(),
      more: response.data.tag_values_more,
      tagMeta,
    };
  }
  if (error) {
    useErrorStore.getState().addError(error);
  }
  return undefined;
}

export function setUpdatedVariable(nameVariable: string | undefined, toggle: boolean) {
  if (nameVariable == null) {
    return;
  }
  useVariableListStore.setState((state) => {
    state.variables[nameVariable] ??= getEmptyVariable();
    state.variables[nameVariable].updated = toggle;
  });
  updateVariables(useStore.getState());
}

export async function getAutoSearchSyncFilter(startIndex: number = 0) {
  const { params, loadMetricsMeta } = useStore.getState();
  await loadAllMeta(params, loadMetricsMeta);
  const { metricsMeta } = useStore.getState();
  const variablesLink: Record<string, VariableParamsLink[]> = {};
  params.plots.forEach(({ metricName }, indexPlot) => {
    const keyPlot = toPlotKey(indexPlot);
    if (metricName === promQLMetric || keyPlot == null) {
      return;
    }
    const meta = metricsMeta[metricName];
    if (!meta) {
      return;
    }
    meta.tags?.forEach((tag, indexTag) => {
      const tagKey = toKeyTag(indexTag);
      if (tagKey && isTagEnabled(meta, tagKey)) {
        const tagName = getTagDescription(meta, indexTag);
        variablesLink[tagName] ??= [];
        variablesLink[tagName].push([keyPlot, tagKey]);
      }
    });
    if (isTagEnabled(meta, TAG_KEY._s)) {
      const tagName = getTagDescription(meta, TAG_KEY._s);
      variablesLink[tagName] ??= [];
      variablesLink[tagName].push([keyPlot, TAG_KEY._s]);
    }
  });
  const addVariables: VariableParams[] = Object.entries(variablesLink)
    .filter(([, link]) => link.length > 1)
    .map(([description, link], index) => {
      const name = isValidVariableName(description)
        ? description
        : `${GET_PARAMS.variableNamePrefix}${startIndex + index}`;
      return {
        ...getEmptyVariableParams(),
        name,
        description: description === name ? '' : description,
        link,
      };
    });
  return addVariables;
}

updateVariables(useStore.getState());
updateTags(useStore.getState());
