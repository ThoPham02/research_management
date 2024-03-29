import { configureStore } from "@reduxjs/toolkit";
import { LoginReducer } from "../pages/common/Login/LoginSlice";
import { DepartmentReducer } from "../components/Shares/Search/Department/DepartmentSlice";
import { FacultyReducer } from "../components/Shares/Search/Faculty/FacultySlice";
import { AdminEventReducer } from "../pages/admin/DashBoard/EventSlice";
import { LectureTopicReducer } from "../pages/lecture/Topic/LectureTopicSlice";
import { CommonEventReducer } from "../components/Shares/Search/Event/CommonEventSlice";
import { CommonTopicReducer } from "../pages/common/Topic/CommonTopicSlice";
import { TopicDetailReducer } from "../components/Shares/Action/Detail/TopicDetailSlice";
import { LectureMyTopicReducer } from "../pages/lecture/MyTopic/LectureMyTopicSlice";
import { DepartmentTopicReducer } from "../pages/department/Topic/DepartmentTopicSlice";
import { FacultyTopicReducer } from "../pages/faculty/Topic/FacultyTopicSlice";
import { StudentMyTopicReducer } from "../pages/student/MyTopic/StudentMyTopicSlice";
import { FacultySubcommitteeReducer } from "../pages/faculty/Report/facultySubcommitteeSlice";
import { CommonSubcommitteeReducer } from "../components/Shares/Search/Subcommittee/CommonSubcommitteeSlice";
import { ItemReducer } from "../components/Shares/Action/AddSubcommittee/itemSlice";
import { ResultReducer } from "../pages/common/Result/ResultSlice";

const store = configureStore({
  reducer: {
    //share reducer
    department: DepartmentReducer,
    faculty: FacultyReducer,
    TopicDetailStore: TopicDetailReducer,


    // common reducer
    login: LoginReducer,
    CommonEventStore: CommonEventReducer,
    CommonTopicStore: CommonTopicReducer,
    ItemStore: ItemReducer,

    CommonSubcommitteeStore:  CommonSubcommitteeReducer,
    TopicMarkStore: ResultReducer,

    // admin reducer
    AdminEventStore: AdminEventReducer,

    //lecture reducer
    LectureTopicStore: LectureTopicReducer,
    LectureMyTopicStore: LectureMyTopicReducer,

    // Department reducer
    DepartmentTopicStore: DepartmentTopicReducer,

    // Faculty reducer 
    FacultyTopicStore: FacultyTopicReducer,
    FacultySubcommitteeStore: FacultySubcommitteeReducer,

    //student reducer
    StudentMyTopicStore: StudentMyTopicReducer,
  },
});

export default store;
