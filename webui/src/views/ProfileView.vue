<script>
import PhotoItem from '../components/PhotoItem.vue';


export default {

    data: function () {
        return {
            errormsg: null,
            loading: false,
            profile: null,
            photos : null,
            //This variable contains the userId that have been passed in the path (id of the profile).
            profileId: this.$route.params.userId,
            //This division is needed because a user can access another user profile so we need both the 
            // id of the user that's visiting and the id of the owner
            //This variable contains the id of the logged user.
            userId: localStorage.getItem("storedData"),
            //This variable contains the token that have been stored after the login.
            token: localStorage.getItem("storedData"),
        };
    },
    methods: {
        async refresh() {
            this.loading = true;
            this.errormsg = null;
            try {
                let response = await this.$axios({
                    method: "get",
                    url: "/user/" + this.profileId + "/profile",
                    headers: {
                        "Content-Type": "application/json",
                    }
                });
                this.profile = response.data;
                this.photos = this.profile.photos;
                document.getElementById("username").textContent = this.profile.username;
            }
            catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },
    },
    mounted() {
        this.$nextTick(function () {
            this.refresh();
        });
    },
    components: { PhotoItem }
}

</script>

<template>
    <div class="row justify-content-between">  

        <!--This column contains all the photos in the stream--> 
        <div class="col-auto text-center" style="padding-left: 12%;">
            <!--Each container is made of the photo and the likes and comment buttons-->	    
            <div>
                <PhotoItem v-for= "photo in photos" :images="photo" :id = "this.userId"></PhotoItem>
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
