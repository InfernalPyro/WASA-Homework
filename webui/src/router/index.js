import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import ProfileView from '../views/ProfileView.vue'
import UploadView from '../views/UploadView.vue'
import LoginView from '../views/LoginView.vue'
import OptionsView from '../views/OptionsView.vue'
import SearchView from '../views/SearchView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: HomeView, name:'home'},
		{path: '/session', component: LoginView, name: 'login'},
		{path: '/user/:userId/profile', component: ProfileView, name: 'profile'},
		{path: '/user/:userId/photo', component: UploadView , name: 'photo'},
		{path: '/user/:userId/options', component: OptionsView , name: 'options'},
		{path: '/search', component: SearchView , name: 'search'},
	]
})

router.beforeEach((to, from, next) => {
	const token = sessionStorage.getItem('storedData')
	// If logged in, or going to the Login page.
	if (token != "null" || to.name === "login" ) {
		// Continue to page.
		next()	

	} else {
		// Not logged in, redirect to login.
		next({name: 'login'})
	}
})


export default router
