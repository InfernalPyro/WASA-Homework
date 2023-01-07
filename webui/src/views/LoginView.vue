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
		}
	},
	methods: {
		async login() {
			this.loading = true;
			this.errormsg = null;
			this.userId = null
			this.username = document.getElementById('username').value
			if(this.username.length < 3 || this.username.length > 16){
				alert("Username must have more than 3 chars and less than 16")
				return
			}
			try {
				let config ={
                    headers:{
                        'Content-Type': "application/json",
                    }
                }
                let data ={
                    "username" : this.username
                }

				let response = await this.$axios.post('/session', data, config);
				this.userId = response.data;
				
				localStorage.setItem('storedData', this.userId)
				this.$emit('localEvent')

				
				this.$router.push("/");
			} catch (e) {
				this.errormsg = e.toString();
			}
			
			this.loading = false;
		},
	},
	mounted() {
		console.log("Local from login " + localStorage.getItem('storedData'))	
	}
}


</script>

<template>
    <h1>Login</h1>
	<label for="username">Username:</label>
  	<input type="text" id="username" name="username" minlenght = "3" maxlenght = "16" pattern = "^[a-zA-Z0-9 ]*$"><br><br>
	<button @click="login()">Login</button>
</template>

<style>
</style>
