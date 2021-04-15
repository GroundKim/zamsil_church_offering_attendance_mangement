import Vue from "vue"
import VueRouter from "vue-router"
import AttendanceInput from "./components/AttendanceInput"
import OfferingInput from "./components/OfferingInput"
import AttendanceAndOfferingView from "./components/AttendanceAndOfferingView"
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
        path: "/attendance-offering/view",
        name: "attendanceAndOfferingView",
        component: AttendanceAndOfferingView
    }
]
})

export default router