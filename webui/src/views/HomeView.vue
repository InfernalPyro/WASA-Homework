<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>
<script>

//This variable contains the id of the user that have been stored after the login. 
//We will use this in the url
//var path = localStorage.getItem('storedData');
//This variable contains the token that have been stored after the login.
var token = localStorage.getItem('storedData');

export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			profile: null,
			path : localStorage.getItem('storedData')
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
                        "Authorization": "Bearer " + token,
                    }
                });
				this.profile = response.data;
                document.getElementById("username").textContent = this.profile.username
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
	},
	mounted() {
        this.$nextTick(function(){
			console.log("Local from Home " + localStorage.getItem('storedData'))
           	
        })
		
	}


}
</script>

<template>


</template>

<style>
</style>
