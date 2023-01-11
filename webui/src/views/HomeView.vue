<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>
<script>


export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			profile: null,
			//This variable contains the id of the user that have been stored after the login. 
			//We will use this in the url
			path: sessionStorage.getItem('storedData'),
			//var path = localStorage.getItem('storedData');
			//This variable contains the token that have been stored after the login.
			token: sessionStorage.getItem('storedData'),
			bannedList: null,
			followList: null,
		}
	},
	methods: {
		async refresh() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios({
                    method:"get",
                    url:'/user/' + this.path + '/profile',
                    headers:{
                        "Content-Type": "application/json",
                        "Authorization": "Bearer " + this.token,
                    }
                });
				this.profile = response.data;
                document.getElementById("username").textContent = this.profile.username

				//Those two session storage will contain both blocked and follow list 
				//to avoid having to make a request each time to get them
				
				this.followList = this.profile.follows;
				this.bannedList = this.profile.banned;
				sessionStorage.setItem('followListData', this.followList);
				sessionStorage.setItem('bannedListData', this.bannedList);


			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
	},
	mounted() {
        this.$nextTick(function(){
            this.refresh() 
        })	
		
	}


}
</script>

<template>


</template>

<style>
</style>
