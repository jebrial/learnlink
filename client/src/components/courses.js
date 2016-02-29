import React from 'react';
import Course from './course';
const Courses = (props) => {
    const courseList = props.courses.map((course) => {
        return (
            <Course
                onCourseSelect={props.onCourseSelect}
                key={course.id}
                    course={course} />
            );
    });
    return (
        <ul className="col-md-4 list-group">
            {courseList}
        </ul>
    );
};

export default Courses;