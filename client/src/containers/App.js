import React, {PropTypes} from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import LearnLinkApp from '../components/LearnLinkApp';
import * as LearnLinkActions from '../actions/learnLinkActions';

class App extends React.Component {
  render() {
    const { learnLinkAppState, actions } = this.props;

    return (
      <LearnLinkApp learnLinkAppState={learnLinkAppState} actions={actions} />
    );
  }
}

App.propTypes = {
  actions: PropTypes.object.isRequired,
  learnLinkAppState: PropTypes.object.isRequired
};

function mapStateToProps(state) {
  return {
    learnLinkAppState: state.learnLinkAppState
  };
}

function mapDispatchToProps(dispatch) {
  return {
    actions: bindActionCreators(LearnLinkActions, dispatch)
  };
}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(App);
