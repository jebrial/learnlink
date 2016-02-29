import React from 'react';

const Course = ({course, onCourseSelect}) => {


    return (
        <li onClick={()=> onCourseSelect(course)}  className="list-group-item">
            {course.name}
        </li>
    );
};

export default Course;