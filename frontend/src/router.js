import Vue from 'vue'
import VueRouter from 'vue-router'
import AttendanceInput from './components/AttendanceInput'
import OfferingInput from './components/OfferingInput'
import OfferingView from './components/OfferingView'
import AttendanceView from './components/AttendanceView'
import MemberManagement from './components/MemberManagement'
import Login from './components/Login'
import AttendanceViewDetail from './components/AttendanceViewDetail'
import OfferingViewDetail from './components/OfferingViewDetail'

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

    {
      path:'/attendance/view/detail',
      name: 'attendanceViewDetail',
      component: AttendanceViewDetail
    },

    {
      path:'/offering/view/detail',
      name: 'offeringViewDetail',
      component: OfferingViewDetail
    },
    
    // redirect "/" to "/login"
    {
      path: '/', redirect: { name: 'login' }
    },

    // simple
    {
      path:'/simple/login',
      name: 'simpleLogin',
      component: Login
    },

    {
      path: '/simple/offering/view',
      name: 'simpleOfferingView',
      component: OfferingView
    },

    {
      path:'/simple/attendance/view',
      name: 'simpleAttendanceView',
      component: AttendanceView
    },

    {
      path:'/simple/attendance/view/detail',
      name: 'simpleAttendanceViewDetail',
      component: AttendanceViewDetail
    },

    {
      path:'/simple/offering/view/detail',
      name: 'simpleOfferingViewDetail',
      component: OfferingViewDetail
    },
]
})

export default router