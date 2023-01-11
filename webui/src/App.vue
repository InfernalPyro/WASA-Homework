<script setup>
import AppComponent from './components/AppComponent.vue';
</script>
<script>


export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			session : sessionStorage.getItem('storedData')
		}
	},
	methods: {
		async refresh() {
			
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
		
	}


}


</script>

<template>
	<AppComponent></AppComponent>
</template>

<style>
</style>
