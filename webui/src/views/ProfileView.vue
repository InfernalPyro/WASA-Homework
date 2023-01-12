<script>
import PhotoItem from '../components/PhotoItem.vue';


export default {

    data: function () {
        return {
            errormsg: null,
            loading: false,
            profile: null,
            photos : null,
            myProfile: null,
            //This variable contains the userId that have been passed in the path (id of the profile).
            profileId: this.$route.params.userId,
            //This division is needed because a user can access another user profile so we need both the 
            // id of the user that's visiting and the id of the owner
            //This variable contains the id of the logged user.
            userId: sessionStorage.getItem("storedData"),
            //This variable contains the token that have been stored after the login.
            token: sessionStorage.getItem("storedData"),
            //Those two contains blocked and follow list of the logged user
            bannedList : sessionStorage.getItem("bannedListData"),
            followList : sessionStorage.getItem("followListData"),
            followFlag : null,
            bannedFlag : null,

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
                        'Authorization': "Bearer " + this.token,
                    }
                });
                this.profile = response.data;
                this.photos = this.profile.photos;
                document.getElementById("username").textContent = this.profile.username; 
                this.followFlag = this.followList.includes(this.profileId);
                this.bannedFlag = this.bannedList.includes(this.profileId);

            }
            catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },
        async banUser() {
            this.loading = true;
            this.errormsg = null;
            try {
                let response = await this.$axios({
                    method: "put",
                    url: "/user/" + this.userId + "/ban/" + this.profileId, 
                    headers: {
                        "Content-Type": "application/json",
                        "Authorization": "Bearer " + this.token,
                    }
                });  
                this.bannedFlag = true;   
                this.followFlag = false;
            }
            catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },
        async unbanUser() {
            this.loading = true;
            this.errormsg = null;
            try {
                let response = await this.$axios({
                    method: "delete",
                    url: "/user/" + this.userId + "/ban/" + this.profileId, 
                    headers: {
                        "Content-Type": "application/json",
                        "Authorization": "Bearer " + this.token,
                    }
                });
                this.bannedFlag = false;
            }
            catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },
        async followUser() {
            this.loading = true;
            this.errormsg = null;
            try {
                let response = await this.$axios({
                    method: "put",
                    url: "/user/" + this.userId + "/follow/" + this.profileId, 
                    headers: {
                        "Content-Type": "application/json",
                        "Authorization": "Bearer " + this.token,
                    }
                });
                this.followFlag = true;
                
            }
            catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },
        async unfollowUser() {
            this.loading = true;
            this.errormsg = null;
            try {
                let response = await this.$axios({
                    method: "delete",
                    url: "/user/" + this.userId + "/follow/" + this.profileId, 
                    headers: {
                        "Content-Type": "application/json",
                        "Authorization": "Bearer " + this.token,
                    }
                });
                this.followFlag = false;
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
            <!--If this profile is not yours there are the follow and block buttons too-->
            <div class="row">
                <div class="col-auto">
                    <button v-if="!loading &&  !this.followFlag && this.profileId != this.userId && !this.bannedFlag" type="button" id = "followButton" style="width: fit-container; margin-left:30%" @click="followUser">
				        Follow
			        </button>
                    <button v-else-if="!loading && this.followFlag && this.profileId != this.userId" type="button" id = "unfollowButton" style="width: fit-container; margin-left:30%" @click="unfollowUser">
				        Unfollow
			        </button>
                </div>
                <div class="col-auto">
                    <button v-if="!loading && !this.bannedFlag && this.profileId != this.userId" type="button" id = "blockButton" style="width: fit-container; margin-left:30%" @click="banUser">
				        Block
			        </button>
                    <button v-else-if="!loading && this.bannedFlag && this.profileId != this.userId" type="button" id = "unblockButton" style="width: fit-container; margin-left:30%" @click="unbanUser">
				        Unblock
			        </button>
                </div>
            </div>

            <div id="rbHiderFillTop"></div>


            <li class="nav-item">
                <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#user-check"/></svg> Following
            </li>
            <li class="nav-item">
                <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#users"/></svg> Followers
            </li>            
            <li class="nav-item">
                <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#slash"/></svg> Blocked
            </li>						
        </div>

    </div>
	
</template>

<style>
</style>
