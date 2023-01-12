<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>
<script>

export default {
	data: function() {
		return {
			//This variable contains the id of the user that have been stored after the login. 
			//We will use this in the url
			path : 0,
			nameView: null,
			userToFind : null,
			//This variable contains the token that have been stored after the login.
			token : sessionStorage.getItem('storedData'),
			session : sessionStorage.getItem('storedData')

		}
	},
    methods:{
		//Function called by the event triggerd on login
        async updatePath(){           
            this.path = sessionStorage.getItem('storedData'); 
			this.nameView = "";
        },
		//Function called when disconnecting 
        async disconnect(){
            sessionStorage.setItem('storedData', null); 
			this.path = null 
			this.nameView = "login";
        },
		async search() {
			this.$router.push("/search");
		},
		
    },
	mounted() {
		 this.$nextTick(function(){
			//By using this check if the browser have just been opened we will start with the loginview
			//(does NOT work on refresh thanks to sessionStorage)
			if (this.session == null){
				this.$router.push("/session");
			}
        })	

		if (sessionStorage.getItem('storedData') != "null"){
			this.path = sessionStorage.getItem('storedData'); 
			this.nameView = "";
		}
		else{
			this.path = 0;
			this.nameView = "login";
		}
	}
}

</script>

<template>

	<!--Show this for every view EXCEPT login-->
	<div v-if="nameView != 'login'">
		<header class="navbar navbar-dark bg-dark flex-md-nowrap p-0 shadow" >
			<div class="col-3">
				<a class="navbar-brand" href="#/">Wasa Photo</a>
			</div>
			<div class="col-2">
				<a class="navbar-brand" id ="username" href="#/">Username</a>
			</div>
			<div class="col-3">
				  <button @click="search">Find other users</button>
			</div>
		</header>

		<div id="rbHiderFillTop"></div>

		<div class="container-fluid">
			<div class="row">
				<nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse" style="height:fit-content;">
					<div class="position-sticky pt-3 ">
						<ul class="nav flex-column">
							<li class="nav-item">
								<RouterLink to="/" class="nav-link">
									<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#home"/></svg>
									Home
								</RouterLink>
							</li>
							<li class="nav-item">
								<RouterLink :to="{ name: 'profile', params: { userId: path }}" class="nav-link">
									<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#user"/></svg>
									Profile
								</RouterLink>
							</li>
							<li class="nav-item">
								<RouterLink :to="{ name: 'photo', params: { userId: path }}" class="nav-link" >
									<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#camera"/></svg>
									Upload
								</RouterLink>
							</li>
							<li class="nav-item">
								<RouterLink :to="{ name: 'options', params: { userId: path }}" class="nav-link">
									<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#settings"/></svg>
									Options
								</RouterLink>
							</li>
							<li class="nav-item" @click="disconnect" >
								<RouterLink to="/session" class="nav-link" >
									<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#log-out"/></svg>
									Disconnect
								</RouterLink>
							</li>
						</ul>
					</div>
				</nav>	
			</div>
		</div>
	</div>

	<main class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
		<RouterView  @localEvent = "updatePath"></RouterView>
	</main>
</template>

<style>
</style>

