import Vue from 'vue'
import VueRouter from 'vue-router'
import AttendanceInput from './components/AttendanceInput'
import OfferingInput from './components/OfferingInput'
import OfferingView from './components/OfferingView'
import AttendanceView from './components/AttendanceView'
import MemberManagement from './components/MemberManagement'
import Login from './components/Login'

Vue.use(VueRouter)

const router = new VueRouter({
    mode: 'history',
    routes: [{
        path: '/attendance/input',
        name: 'attendanceInput',
        component: AttendanceInput
    },

    {
        path: '/offering/input',
        name: 'offeringInput',
        component: OfferingInput
    },

    {
        path: '/offering/view',
        name: 'offeringView',
        component: OfferingView
    },

    {
        path:'/attendance/view',
        name: 'attendanceView',
        component: AttendanceView
    },

    {
        path:'/login',
        name: 'login',
        component: Login
    },

    {
        path:'/member/management',
        name:'memberManagement',
        component: MemberManagement
    },
]
})

export default router