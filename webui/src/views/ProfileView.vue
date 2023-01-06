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
		async refresh() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios({
                    method:"get",
                    url:'/user/' + 1 + '/profile',
                    headers:{
                        "Content-Type": "application/json",
                        "Authorization": "Bearer 1",
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
           this.refresh() 
        })
		
	}
}

</script>

<template>
    <div class="row justify-content-between">


        <!--This column contains all the photos in the stream--> 
        <div class="col-auto">
            <!--Each container is made of the photo and the likes and comment buttons-->
		    <div class="container">
                <!--First row contains the image-->
                <div class="row">
                    <!--<img src = "data:image/gif;base64,........">-->
                </div>
                <!--Second row contains likes and comment-->
                <div class="row justify-content-around">
                    <div class="col">
                        <li class="nav-item">
                            <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#heart"/></svg> 
                        </li>
                    </div>
                    <div class="col">
                        <li class="nav-item">
                            <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#message-square"/></svg> 
                        </li>
                    </div>                   
                </div>
            </div>


	    </div>


        <!--This is a fixed in page column that contains the profile informations-->
        <div class="col-3" style="padding-right: 100px; border-left-style: solid;">
            <li class="nav-item">
                <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#user-check"/></svg> Following
            </li>
            <li class="nav-item">
                <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#users"/></svg> Followers
            </li>            
            <li class="nav-item">
                <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#slash"/></svg> Ban
            </li>						
        </div>

    </div>
	
</template>

<style>
</style>
