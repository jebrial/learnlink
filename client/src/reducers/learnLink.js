import {SELECT_LEARNLINK, CHANGE_NAME, CHANGE_URL, CHANGE_NOTE } from '../constants/ActionTypes';

const initialState = {
  learnLinks: [
    { "id": 11, "name": "Redux Tutorial", "url": "http://redux.js.org/docs/introduction/", "note": "A great starting point to learn Redux", "priority": 1, "interval":89, "checkedOnDate":"date here", "isDone": false},
    { "id": 12, "name": "JWT for react/redux" , "url": "https://auth0.com/blog/2016/01/04/secure-your-react-and-redux-app-with-jwt-authentication/", "note": "next read!", "priority": 1, "interval":89, "checkedOnDate":"date here", "isDone": false},
    { "id": 13, "name": "JS interview questions" , "url": "http://www.thatjsdude.com/interview/js2.html", "note": "Very Helpful", "priority": 1, "interval":89, "checkedOnDate":"date here", "isDone": false},
    { "id": 14, "name": "Stanford Distributed Systems Videos" , "url": "https://www.youtube.com/playlist?list=PL72C36006AD9CED5C", "note": "Interesting watch", "priority": 1, "interval":89, "checkedOnDate":"date here", "isDone": false},
    { "id": 15, "name": "A software Dev's Reading List" , "url": "http://stevewedig.com/2014/02/03/software-developers-reading-list/", "note": "Endless study", "priority": 1, "interval":89, "checkedOnDate":"date here", "isDone": false},
    { "id": 16, "name": "Tour of Go" , "url": "https://tour.golang.org/welcome/1", "note": "Fun class", "priority": 1, "interval":89, "checkedOnDate":"date here", "isDone": false},
    { "id": 17, "name": "Linear Algebra On Khan Academy" , "url": "https://www.khanacademy.org/math/linear-algebra", "note": "still learning", "priority": 1, "interval":89, "checkedOnDate":"date here", "isDone": false},
    { "id": 18, "name": "Coursera Algorithms" , "url": "https://www.coursera.org/course/algo", "note": "Always need to review this", "priority": 1, "interval":89, "checkedOnDate":"date here", "isDone": false},
    { "id": 19, "name": "Node School" , "url": "http://nodeschool.io/#workshoppers", "note": "Great Node starter", "priority": 1, "interval":89, "checkedOnDate":"date here", "isDone": false},
    { "id": 20, "name": "Think Big" , "url": "http://bigthink.com/", "note": "For inspiration", "priority": 1, "interval":89, "checkedOnDate":"date here", "isDone": false}
  ],
  selectedLearnLink: null

};

export default function learnLinkAppState(state = initialState, action) {
  let newState = null;

  switch (action.type) {
    case SELECT_LEARNLINK:
          newState = Object.assign({}, state);
          newState.selectedLearnLink = state.learnLinks.find(learnLink =>
            learnLink.id = action.id
          );
          return newState;
    case CHANGE_NAME:
        newState = Object.assign({}, state);
        newState.heroes = newState.learnLinks.map(learnLink => {
            if (learnLink.id === action.id) {
              learnLink.name = action.newName;
            }
            return learnLink;
            }
        );
        return newState;
    case CHANGE_URL:
        newState = Object.assign({}, state);
        newState.learnLinks = newState.learnLinks.map(learnLink => {
            if (learnLink.id === action.id) {
              learnLink.url = action.newUrl;
            }
            return learnLink;
            }
        );
        return newState;
    case CHANGE_NOTE:
        newState = Object.assign({}, state);
        newState.learnLinks = newState.learnLinks.map(learnLink => {
            if (learnLink.id === action.id) {
              learnLink.note = action.newNote;
            }
              return learnLink;
            }
        );
        return newState;
    default:
      return state;
  }
}
