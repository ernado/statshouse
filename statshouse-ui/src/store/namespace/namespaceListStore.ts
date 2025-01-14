import { createStore } from '../createStore';
import { useErrorStore } from '../errors';
import { GET_PARAMS } from '../../api/enum';
import {
  apiNamespaceFetch,
  ApiNamespacePost,
  ApiNamespacePut,
  apiPostNamespaceFetch,
  apiPutNamespaceFetch,
} from '../../api/namespace';
import { apiNamespaceListFetch, NamespaceShort } from '../../api/namespaceList';
import { sortByKey } from '../../view/utils';

export const namespaceListErrors = 'groupListErrors';

export type NamespaceListStore = {
  list: NamespaceShort[];
};
export const useNamespaceListStore = createStore<NamespaceListStore>(() => ({
  list: [],
}));

let loadListErrorRemover: () => void;

export async function namespaceListLoad() {
  loadListErrorRemover?.();
  const { response, error } = await apiNamespaceListFetch();
  if (error) {
    loadListErrorRemover = useErrorStore.getState().addError(error, namespaceListErrors);
  }
  if (response) {
    useNamespaceListStore.setState((s) => {
      const list = response.data.namespaces ?? [];
      list.sort(sortByKey.bind(undefined, 'name'));
      if (!list.some((n) => n.id <= 0)) {
        list.unshift({ id: -9999, name: 'default', weight: 1 });
      }
      s.list = list;
    });
  }
}

let loadErrorRemover: () => void;

export async function namespaceLoad(id: number) {
  loadErrorRemover?.();
  const { response, error } = await apiNamespaceFetch({ [GET_PARAMS.metricsNamespacesID]: id.toString() });
  if (error) {
    loadErrorRemover = useErrorStore.getState().addError(error, namespaceListErrors);
  }
  if (response) {
    return response.data;
  }
  return null;
}

let addErrorRemover: () => void;

export async function namespaceAdd(namespace: ApiNamespacePut) {
  addErrorRemover?.();
  const { response, error } = await apiPutNamespaceFetch(namespace);
  if (error) {
    addErrorRemover = useErrorStore.getState().addError(error, namespaceListErrors);
  }
  await namespaceListLoad();
  if (response) {
    return response.data;
  }
  return null;
}
let saveErrorRemover: () => void;
export async function namespaceSave(namespace: ApiNamespacePost) {
  saveErrorRemover?.();
  const { response, error } = await apiPostNamespaceFetch(namespace);
  if (error) {
    saveErrorRemover = useErrorStore.getState().addError(error, namespaceListErrors);
  }
  await namespaceListLoad();
  if (response) {
    return response.data;
  }
  return null;
}
