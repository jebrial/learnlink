import React, {Component} from 'react';

export default class AddCourse extends Component {
    render() {
        return (
            <div>
                <h2>Add New Course:</h2>
                <div>
                    <label>name:</label>
                    <input value="" placeholder="name" />
                </div>
                <div>
                    <label>url:</label>
                    <input value=""  placeholder="url" />
                </div>
                <div>
                    <label>note:</label>
                    <input value="" placeholder="note" />
                </div>
            </div>
        );
    }
}