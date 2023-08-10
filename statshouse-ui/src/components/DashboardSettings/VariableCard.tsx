// Copyright 2023 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

import produce from 'immer';
import { getMetricFullName, isValidVariableName } from '../../view/utils';
import React, { useCallback, useEffect, useState } from 'react';
import { PlotStore } from '../../store';
import { MetricMetaValue } from '../../api/metric';
import { isNil, isNotNil } from '../../common/helpers';
import { VariablePlotLinkSelect } from './VariablePlotLinkSelect';
import { ReactComponent as SVGTrash } from 'bootstrap-icons/icons/trash.svg';
import { ReactComponent as SVGChevronUp } from 'bootstrap-icons/icons/chevron-up.svg';
import { ReactComponent as SVGChevronDown } from 'bootstrap-icons/icons/chevron-down.svg';
import { ToggleButton } from '../UI';
import cn from 'classnames';
import { PlotParams, VariableParams } from '../../url/queryParams';

export type VariableCardProps = {
  indexVariable: number;
  variable?: VariableParams;
  setVariable?: (indexVariable: number, value?: React.SetStateAction<VariableParams>) => void;
  plots: PlotParams[];
  plotsData: PlotStore[];
  metricsMeta: Record<string, MetricMetaValue>;
};

export function VariableCard({
  indexVariable,
  variable,
  setVariable,
  plots,
  plotsData,
  metricsMeta,
}: VariableCardProps) {
  const [open, setOpen] = useState(false);
  const [valid, setValid] = useState(false);

  useEffect(() => {
    setValid(isValidVariableName(variable?.name ?? `v${indexVariable}`));
  }, [indexVariable, variable?.name]);

  const setName = useCallback(
    (e: React.FormEvent<HTMLInputElement>) => {
      const value = e.currentTarget.value;
      const valid = isValidVariableName(value);
      if (!valid && value !== '') {
        e.preventDefault();
        return;
      }
      setVariable?.(
        indexVariable,
        produce((v) => {
          if (value === '') {
            v.name = `v${indexVariable}`;
          } else {
            v.name = value;
          }
        })
      );
    },
    [indexVariable, setVariable]
  );

  const setDescription = useCallback(
    (e: React.FormEvent<HTMLInputElement>) => {
      const value = e.currentTarget.value;
      setVariable?.(
        indexVariable,
        produce((v) => {
          v.description = value;
        })
      );
    },
    [indexVariable, setVariable]
  );
  const remove = useCallback(() => {
    setVariable?.(indexVariable, undefined);
  }, [indexVariable, setVariable]);

  const plotLink = useCallback(
    (indexPlot: number, selectTag?: number) => {
      setVariable?.(
        indexVariable,
        produce((v) => {
          const indexLink = v.link.findIndex(([iPlot]) => indexPlot === iPlot);
          if (indexLink > -1) {
            if (isNil(selectTag)) {
              v.link.splice(indexLink, 1);
            } else {
              v.link[indexLink] = [indexPlot, selectTag];
            }
          } else if (isNotNil(selectTag)) {
            v.link.push([indexPlot, selectTag]);
          }
          if (v.link.length === 0) {
            v.args = { groupBy: false, negative: false };
          }
        })
      );
    },
    [indexVariable, setVariable]
  );

  if (!variable) {
    return null;
  }
  return (
    <div className="card">
      <div className="card-body">
        <div className="d-flex align-items-center">
          <div className="input-group">
            <ToggleButton
              className="btn btn-outline-primary rounded-start"
              checked={open}
              onChange={setOpen}
              title={open ? 'collapse' : 'expand'}
            >
              {open ? <SVGChevronUp /> : <SVGChevronDown />}
            </ToggleButton>
            <span className="input-group-text">
              {variable.link.length} of {plots.length}
            </span>
            <input
              name="name"
              className={cn('form-control', !valid && 'border-danger')}
              placeholder={`v${indexVariable}`}
              value={variable.name !== `v${indexVariable}` ? variable.name : ''}
              onInput={setName}
            />
            <input
              name="description"
              className="form-control"
              placeholder="description"
              value={variable.description}
              onInput={setDescription}
            />
          </div>
          <button className="btn btn-outline-danger ms-2" onClick={remove} title="Remove">
            <SVGTrash />
          </button>
        </div>
        {!valid && (
          <div className="text-danger small">Not valid variable name, can content only 'a-z', '0-9' and '_' symbol</div>
        )}
        {open && (
          <div>
            <table className="table align-middle table-borderless">
              <tbody className="mb-2 border-bottom-1">
                {plots.map((plot, indexPlot) => (
                  <tr key={indexPlot}>
                    <td className="text-end pb-0 ps-0">{getMetricFullName(plot, plotsData[indexPlot])}</td>
                    <td className="pb-0 pe-0">
                      <VariablePlotLinkSelect
                        indexPlot={indexPlot}
                        selectTag={variable.link.find(([p]) => p === indexPlot)?.[1] ?? undefined}
                        metricMeta={metricsMeta[plot.metricName]}
                        onChange={plotLink}
                      />
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}
      </div>
    </div>
  );
}