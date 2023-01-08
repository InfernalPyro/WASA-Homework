import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import ProfileView from '../views/ProfileView.vue'
import UploadView from '../views/UploadView.vue'
import LoginView from '../views/LoginView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: HomeView, name:'home', meta:{requiresAuth: true}},
		{path: '/session', component: LoginView, name: 'login'},
		{path: '/user/:userId/profile', component: ProfileView, name: 'profile', meta:{requiresAuth: true}},
		{path: '/user/:userId/photo', component: UploadView , name: 'photo', meta:{requiresAuth: true}},
	]
})


var logged = localStorage.getItem('storedData');
router.beforeEach((to, from, next) => {
	if (to.matched.some(record => record.meta.requiresAuth)) {
	  // this route requires auth, check if logged in
	  // if not, redirect to login page.
	  if (logged === null || logged === "") {
		next({ name: 'login'})
	  } else {
		next() // go to wherever I'm going
	  }
	} else {
	  next() // does not require auth, make sure to always call next()!
	}
  })

export default router
