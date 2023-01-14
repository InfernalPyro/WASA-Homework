<script>

var base64String = null;

export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			profile: null,
            //This variable contains the userId that have been passed in the path.
            userId: this.$route.params.userId,
            //This variable contains the token that have been stored after the login.
			token : sessionStorage.getItem('storedData'),
            
		}
	},
	methods: {
		async refresh() {
			this.loading = true;
			this.errormsg = null;
			try {
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
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},

        async uploadPhoto() {
			this.loading = true;
			this.errormsg = null;
			try {
                let config ={
                    headers:{
                        'Content-Type': "application/json",
                        "Authorization": "Bearer " + this.token,
                    }
                }
                let data ={
                    "image" : base64String
                }

				let response = await this.$axios.post('/user/'+ this.userId + '/photo/', data, config);
                const container = document.querySelector("#imageContainer");
                while (container.firstChild) {
                    container.removeChild(container.lastChild);
                }
				this.$router.push("/user/"+ this.userId +"/profile");
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},

        async convertPhoto(){
            var file = document.getElementById('photo').files[0];
            var reader  = new FileReader();
            reader.onload = function(e)  {
                /*Show the photo on screen before uploading it*/ 
                var image = document.createElement("img");
                image.src = e.target.result;
                const container = document.querySelector("#imageContainer");
                while (container.firstChild) {
                    container.removeChild(container.lastChild);
                }
                container.appendChild(image);
                
                /*And convert in base64*/ 
                /*This will contain the base64 string*/
                base64String = reader.result.replace("data:", "").replace(/^.+,/, "");
            }
            reader.readAsDataURL(file);
            document.getElementById("uploadButton").disabled = false;

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

        <LoadingSpinner v-if="loading"></LoadingSpinner>

        <!--This column contains all the photos in the stream--> 
        <div class="col-8">
            <h1 class="center">Upload a new photo</h1>
            <div id="rbHiderFillTop"></div>

            <div class="row">     
                <div class="container" id="imageContainer" style="width: 50%; height : 50%; border:hidden">
                </div>
                
            </div>
            
            <div id="rbHiderFillTop"></div>

            <div class="row">
                <input type="file" accept="image/*" id="photo" @change="convertPhoto">
            </div>
	    </div>

        <!--This is a fixed in page column that contains the profile informations-->
        <div class="col-3" style="padding-right: 100px; border-left-style: solid;">
             <div class="row">
                <button v-if="!loading" type="button" id = "uploadButton" class="btn btn-primary" style="width: fit-container; margin-left:30%" @click="uploadPhoto" disabled>
				    Upload
			    </button>
            </div>
        </div>


    </div>
	
</template>

<style>
</style>