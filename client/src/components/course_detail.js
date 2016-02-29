import React from 'react';

const CourseDetail = ({course}) => {
    if(!course){ return <div>Loading...</div>;}

    return (
        <div className="course-detail col-md-8">
            <div className="details">
                <div>{course.name}</div>
                <div>{course.url}</div>
                <div>{course.note}</div>
            </div>
        </div>
    )
};

export default CourseDetail;