import Vue from "vue"
import VueRouter from "vue-router"
import Teacher from "./components/Teacher"
Vue.use(VueRouter)

const router = new VueRouter({
    mode: "history",
    routes: [{
        path: "/input",
        name: "input",
        component: Teacher
    }]
})

export default router