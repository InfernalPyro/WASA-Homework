<script>
import PhotoItem from '../components/PhotoItem.vue';
import UsersModal from '../components/UsersModal.vue';



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
            usersList : null,
            isModalVisible : false,
            modalType :null,
            youAreBannedFlag : false,
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
                
                if (this.profile.banned != null){
                    if (this.profile.banned.includes(Number(this.token))){
                        this.youAreBannedFlag = true;
                    }
                }
                

                if (this.profileId == this.token){
                    this.followList = this.profile.follows;
				    this.bannedList = this.profile.banned;
				    sessionStorage.setItem('followListData', this.followList);
				    sessionStorage.setItem('bannedListData', this.bannedList);
                }
                
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

        async openFollowingModal() {
            this.isModalVisible = true;
            this.usersList = this.profile.follows
            this.modalType = "Following"
            document.getElementsByTagName("body")[0].style = "overflow-y: hidden; "
        },
        async openFollowersModal() {
            this.isModalVisible = true;
            this.usersList = this.profile.followers 
            this.modalType = "Followers"
            document.getElementsByTagName("body")[0].style = "overflow-y: hidden; "
        },
        async openBlockedModal() {
            this.isModalVisible = true;
            this.usersList = this.profile.banned 
            this.modalType = "Blocked"
            document.getElementsByTagName("body")[0].style = "overflow-y: hidden; "
        },
        async closeModal() {
            this.isModalVisible = false;
            document.getElementsByTagName("body")[0].style = "overflow-y: scroll;"
        },
    },
    mounted() {
        this.$nextTick(function () {
            this.refresh();
        });
    },
    beforeRouteUpdate(to, from, next) {
        // Call the API query method when the URL changes
        this.closeModal();
        next()
        
       
    },
    components: { PhotoItem , UsersModal}
}

</script>

<template>

    <template  v-if="!youAreBannedFlag">
        <div  class="row justify-content-between">  
            <UsersModal v-if="isModalVisible" :list="this.usersList" :type="this.modalType" @close="closeModal"/>

            <!--This column contains all the photos in the stream--> 
            <div class="col-auto text-center" style="padding-left: 12%;">
                <!--Each container is made of the photo and the likes and comment buttons-->	    
                <div>
                    <PhotoItem v-for= "photo in photos" :images="photo" :id = "this.userId" :key="photo.photoId" @deleted="refresh"></PhotoItem>
                </div>
	        </div>


            <!--This is a fixed in page column that contains the profile informations-->
            <div class="col-3" style="padding-right: 100px; border-left-style: solid;">
                <!--If this profile is not yours there are the follow and block buttons too-->
                <div class="row">
                    <div class="col-auto">
                        <template v-if="!loading &&  !this.followFlag && this.profileId != this.userId && !this.bannedFlag">
                            <button type="button" id = "followButton" style="width: fit-container; margin-left:30%" @click="followUser">
		    		            Follow
		    	            </button>
                        </template>
                        <template v-else-if="!loading && this.followFlag && this.profileId != this.userId">
                            <button type="button" id = "unfollowButton" style="width: fit-container; margin-left:30%" @click="unfollowUser">
		    		            Unfollow
		    	            </button>
                        </template>
                    
                    </div>
                    <div class="col-auto">
                        <template v-if="!loading && !this.bannedFlag && this.profileId != this.userId">
                            <button type="button" id = "blockButton" style="width: fit-container; margin-left:30%" @click="banUser">
		    		            Block
		    	            </button>
                        </template>
                        <template v-else-if="!loading && this.bannedFlag && this.profileId != this.userId">
                            <button type="button" id = "unblockButton" style="width: fit-container; margin-left:30%" @click="unbanUser">
			    	            Unblock
			                </button>
                        </template>   
                    </div>
                </div>

                <div id="rbHiderFillTop"></div>


                <li class="nav-item">
                    <svg class="feather" @click="openFollowingModal"><use href="/feather-sprite-v4.29.0.svg#user-check"/></svg> Following
                </li>
                <li class="nav-item">
                   <svg class="feather" @click="openFollowersModal"><use href="/feather-sprite-v4.29.0.svg#users"/></svg> Followers
                </li>            
                <li class="nav-item">
                    <svg class="feather" @click="openBlockedModal"><use href="/feather-sprite-v4.29.0.svg#slash"/></svg> Blocked
                </li>						
            </div>
        </div>

    </template>
    

    <div v-else>
        <h4>You are banned from this profile  </h4>
    </div>
	
</template>

<style>
</style>
