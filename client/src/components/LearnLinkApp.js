import React, {PropTypes} from 'react';
import LearnLinkDetail from './LearnLinkDetail';
class LearnLinkApp extends React.Component {
  constructor(props) {
    super(props);
  }

  onLearnLinkClick(id) {
    this.props.actions.selectLearnLink(id);
  }

  render() {
    let appState = this.props.learnLinkAppState;
    let actions = this.props.actions;

    return (
      <div>
        <h2>Study List</h2>
        <ul className="learnLinks">
          {appState.learnLinks.map(learnLink =>
            <li className={appState.selectedLearnLink && learnLink.id === appState.selectedLearnLink.id ? 'selected' : ''}
              key={learnLink.id}
              onClick={this.onLearnLinkClick.bind(this, learnLink.id)}>
              <span>{learnLink.name}</span>
              <span>{learnLink.url}</span>
              <span>{learnLink.note}</span>
            </li>
          )}
        </ul>
        {appState.selectedLearnLink ?
          <LearnLinkDetail key={appState.selectedLearnLink.id} learnLink={appState.selectedLearnLink} actions={actions} />
          : null
        }
      </div>
    );
  }
}

LearnLinkApp.propTypes = {
  actions: PropTypes.object.isRequired,
  learnLinkAppState: PropTypes.object.isRequired
};

export default LearnLinkApp;
