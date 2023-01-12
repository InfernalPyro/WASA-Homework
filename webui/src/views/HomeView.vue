<script setup>
import PhotoItem from '../components/PhotoItem.vue';

</script>
<script>


export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			profile: null,
			stream : null,
			//This variable contains the id of the user that have been stored after the login. 
			//We will use this in the url
			userId: sessionStorage.getItem('storedData'),
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
				//Query to get the profile
				let response = await this.$axios({
                    method:"get",
                    url:'/user/' + this.userId + '/profile',
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
		async getStream() {
			this.loading = true;
			this.errormsg = null;
			try {
				//Query to get the profile
				let response = await this.$axios({
                    method:"get",
                    url:'/user/' + this.userId + '/home',
                    headers:{
                        "Content-Type": "application/json",
                        "Authorization": "Bearer " + this.token,
                    }
                });
				this.stream = response.data;

			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
	},
	mounted() {
        this.$nextTick(function(){
            this.refresh() 
			this.getStream() 
        })		
	},
	components: { PhotoItem }


}
</script>

<template>
		<!--This column contains all the photos in the stream--> 
        <div class="col-8 text-center" style="padding-left: 12%;">
            <!--Each container is made of the photo and the likes and comment buttons-->	    
            <div>
                <PhotoItem v-for= "photo in stream" :images="photo" :id = "userId" :key="photo.photoId"></PhotoItem>
            </div>
	    </div>

</template>

<style>
</style>
