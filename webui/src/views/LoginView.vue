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
			userId: null,
			username: null,
		}
	},
	methods: {
		async login() {
			this.loading = true;
			this.errormsg = null;

			//Check on the name just entered
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
				document.getElementById('username').value = "";
				sessionStorage.setItem('storedData', this.userId);
				this.$emit('localEvent');

				this.$router.push("/");
			} catch (e) {
				this.errormsg = e.toString();
			}
			
			this.loading = false;
		},
		
	},
	emits: 
		["localEvent"], 

	mounted() {
		sessionStorage.setItem('storedData', null);
	}
}


</script>

<template>
	<div id="rbHiderFillTop"></div>
	<div id="rbHiderFillTop"></div>
	<div class="row">
		<div class ="col-auto" style="padding-left: 20vw">
			<div class="container" style="border-radius: 20px;">
				<h1>Login</h1>
				<div id="rbHiderFillTop"></div>
				<label for="username">Username:</label>
  				<input type="text" id="username" name="username" minlenght = "3" maxlenght = "16" pattern = "^[a-zA-Z0-9 ]*$"><br><br>
				<br>
				<button @click="login">Login</button>
				<br>
			</div>
		</div>
	</div>
    
</template>

<style>
</style>
