<script>

var base64String = "";


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

        uploadPhoto: async function () {
			this.loading = true;
			this.errormsg = null;

			try {
                let config ={
                    headers:{
                        'Content-Type': "application/json",
                        "Authorization": "Bearer 1",
                    }
                }
                let data ={
                    "image" : base64String
                }

				let response = await this.$axios.post('/user/'+ 1 + '/photo/', data, config);
				this.$router.push("/profile");
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		}

	},
	mounted() {
        this.$nextTick(function(){
           this.refresh() 

        document.getElementById('photo').addEventListener('change', function() {
            /*Show the photo on screen before uploading it*/ 
            var file = document.getElementById('photo').files[0];
            var reader  = new FileReader();
            reader.onload = function(e)  {
                var image = document.createElement("img");
                image.src = e.target.result;
                const container = document.querySelector("#imageContainer");
                while (container.firstChild) {
                    container.removeChild(container.lastChild);
                }
                container.appendChild(image);
            }
            reader.readAsDataURL(file);

            /*Convert in base64*/ 
            var file = document.querySelector(
                'input[type=file]')['files'][0];
     
            var reader = new FileReader();         
            reader.onload = function () {
                /*This will contain the base64 string*/
                base64String = reader.result.replace("data:", "")
                    .replace(/^.+,/, "");
            
            }       
            reader.readAsDataURL(file);

            document.getElementById("uploadButton").disabled = false;
            
        })


        })		
	} 
}



</script>

<template>
    <div class="row justify-content-between">


        <!--This column contains all the photos in the stream--> 
        <div class="col-8">
            <div class="row">
                <h1 class="center">Upload a new photo</h1>
                <div id="rbHiderFillTop"></div>
                <input type="file" accept="image/*" id="photo">
            </div>
            
            <div id="rbHiderFillTop"></div>
            <div class="row">
                <div class="container" id="imageContainer" style="width: 50%; height : 50%;">

                </div>
            </div>
           
			    
			<LoadingSpinner v-if="loading"></LoadingSpinner>
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