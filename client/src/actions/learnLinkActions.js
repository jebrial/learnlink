import * as types from '../constants/ActionTypes';

export function changeName(id, newName) {
  return { type: types.CHANGE_NAME, id, newName };
}

export function changeUrl(id, newUrl) {
  return { type: types.CHANGE_URL, id, newUrl };
}

export function changeNote(id, newNote) {
  return { type: types.CHANGE_NOTE, id, newNote };
}

export function selectLearnLink(id) {
  return { type: types.SELECT_LEARNLINK, id };
}
