import React, { Component } from 'react';
import ReactDOM from 'react-dom';
import Courses from './components/courses';
import CourseDetail from './components/course_detail';
import AddCourse from './components/add_course';

class App extends Component {
    constructor(props) {
        super(props);
        this.state = {
            courses: [],
            selectedCourse: null
        };
    }

    componentDidMount() {
        this.loadInitialCourseList()

    }

    loadInitialCourseList() {
        //Api call here

        // ({token: token, term: term}, courses => {
        //    this.setState({
        //        courses :courses,
        //        selectedCourse: courses[0]
        //    });
        //});
        const courses = [
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
        ];

        this.setState({
            courses: courses,
            selectedCourse: courses[0]
        });

    }
    render() {
        return (
            <div>
                <CourseDetail course={this.state.selectedCourse} />
                <AddCourse />
                <Courses
                    onCourseSelect={ selectedCourse => this.setState({selectedCourse}) }
                    courses={this.state.courses} />
            </div>
        );
    }
}

ReactDOM.render(<App />, document.querySelector('.container'));