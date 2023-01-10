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

router.beforeEach((to, from, next) => {
	const token = localStorage.getItem('storedData')
	// If logged in, or going to the Login page.
	if (token == null || to.name === 'login') {
		// Continue to page.
		next()
	} else {
		// Not logged in, redirect to login.
		next({name: 'login'})
	}
})


export default router
