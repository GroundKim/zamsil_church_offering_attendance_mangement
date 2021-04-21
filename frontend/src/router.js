import Vue from "vue"
import VueRouter from "vue-router"
import AttendanceInput from "./components/AttendanceInput"
import OfferingInput from "./components/OfferingInput"
import OfferingView from "./components/OfferingView"
import AttendanceView from "./components/AttendanceView"

Vue.use(VueRouter)

const router = new VueRouter({
    mode: "history",
    routes: [{
        path: "/attendance/input",
        name: "attendanceInput",
        component: AttendanceInput
    },

    {
        path: "/offering/input",
        name: "offeringInput",
        component: OfferingInput
    },

    {
        path: "/offering/view",
        name: "offeringView",
        component: OfferingView
    },

    {
        path:"/attendance/view",
        name: "attendanceView",
        component: AttendanceView
    }
]
})

export default router