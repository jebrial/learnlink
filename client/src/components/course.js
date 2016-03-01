import React from 'react';
import Card from 'material-ui/lib/card/card';
import CardActions from 'material-ui/lib/card/card-actions';
import CardHeader from 'material-ui/lib/card/card-header';
import FlatButton from 'material-ui/lib/flat-button';
import CardText from 'material-ui/lib/card/card-text';

const Course = ({course}) => {


    return (
            <Card>
                <CardHeader
                    title={course.name}
                    subtitle=""
                    actAsExpander={true}
                    showExpandableButton={true}
                />
                <CardText expandable={true}>
                    {course.note}
                </CardText>
                <CardActions expandable={true}>
                    <FlatButton label="Edit"/>
                    <FlatButton label="Checkoff" />
                    <FlatButton label="Finish"/>
                </CardActions>
            </Card>
    );
};

export default Course;