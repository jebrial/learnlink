import React, {PropTypes} from 'react';

class LearnLinkDetail extends React.Component {
  constructor(props) {
    super(props);
  }

  onLearnLinkNameChange(e){
    this.props.actions.changeName(this.props.learnLink.id, e.target.value);
  }

  onLearnLinkUrlChange(e) {
    this.props.actions.changeUrl(this.props.learnLink.id, e.target.value);
  }

  onLearnLinkNoteChange(e) {
    this.props.actions.changeNote(this.props.learnLink.id, e.target.value);
  }

  render() {
    let learnLink = this.props.learnLink;

    return (
      <div>
        <h2>{learnLink.name} Edit</h2>
        <div>
          <label>name:</label>
          <input value={learnLink.name} onChange={this.onLearnLinkNameChange.bind(this)} placeholder="name" />
        </div>
        <div>
          <label>url:</label>
          <input value={learnLink.url} onChange={this.onLearnLinkUrlChange.bind(this)} placeholder="url" />
        </div>
        <div>
          <label>note:</label>
          <input value={learnLink.note} onChange={this.onLearnLinkNoteChange.bind(this)} placeholder="note" />
        </div>
      </div>
    );
  }
}

LearnLinkDetail.propTypes = {
  learnLink: PropTypes.object.isRequired,
  actions: PropTypes.object.isRequired
};

export default LearnLinkDetail;
