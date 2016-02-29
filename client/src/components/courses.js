import React from 'react';
import Course from './course';
const Courses = (props) => {
    if(props.courses.length < 1){ return <div>Loading...</div>;}
    const courseList = props.courses.map((course) => {
        return (
            <Course
                key={course.id}
                    course={course} />
            );
    });
    return (
        <div>
            {courseList}
        </div>
    );
};

export default Courses;