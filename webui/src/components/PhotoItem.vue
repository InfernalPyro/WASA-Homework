<script>
import CommentModal from './CommentModal.vue';


export default {
    props: ['images','id'],
    data: function() {
		return {
            photos: this.images,
            myId : this.id,
            token: sessionStorage.getItem("storedData"),
            likeCount: null,
            likeFlag: null,
            commentsCount: null,
            isModalVisible : false,
        }
    },
    methods:{
        async likePhoto() {
            this.loading = true;
            this.errormsg = null;
            try {
                
                let response = await this.$axios({
                    method: "put",
                    url: "/profile/" + this.photos.profileId + "/photo/" + this.photos.photoId + "/likes/" + this.myId,
                    headers: {
                        "Content-Type": "application/json",
                        "Authorization": "Bearer " + this.token,
                    }
                });
                this.likeCount ++;
                this.likeFlag = true;
            }
            catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },
        async unlikePhoto() {
            this.loading = true;
            this.errormsg = null;
            try {
                
                let response = await this.$axios({
                    method: "delete",
                    url: "/profile/" + this.photos.profileId + "/photo/" + this.photos.photoId + "/likes/" + this.myId,
                    headers: {
                        "Content-Type": "application/json",
                        "Authorization": "Bearer " + this.token,
                    }
                });
                this.likeCount --;
                this.likeFlag = false;
            }
            catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },
        async openCommentModal() {
            this.isModalVisible = true;
            document.getElementsByTagName("body")[0].style = "overflow-y: hidden; "

        },
        async closeModal() {
            this.isModalVisible = false;
            document.getElementsByTagName("body")[0].style = "overflow-y: scroll;"
        },
        async blockUser() {
            this.loading = true;
            this.errormsg = null;
            try {
                let response = await this.$axios({
                    method: "put",
                    url: "/user/" + this.myId + "/ban/" + this.photos.profileId, 
                    headers: {
                        "Content-Type": "application/json",
                        "Authorization": "Bearer " + this.token,
                    }
                });
                this.$router.push("/");
            }
            catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },
        async deletePhoto() {
            this.loading = true;
            this.errormsg = null;
            try {
                let response = await this.$axios({
                    method: "delete",
                    url: "/user/" + this.myId + "/photo/" + this.photos.photoId,
                    headers: {
                        "Content-Type": "application/json",
                        "Authorization": "Bearer " + this.token,
                    }
                });
            this.$emit("deleted")
            }
            catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },
        async goToProfile(){
            this.$router.push("/user/" + this.photos.profileId + "/profile");
        },
    },
    
    created() {
        this.$nextTick(function () {
            if (this.photos.likes != null){
                this.likeCount =  this.photos.likes.length;
                this.likeFlag = this.photos.likes.includes(Number(this.myId))
            }
            else{
                this.likeCount = 0;
                this.likeFlag = false;
            }

            this.commentsCount = (this.photos.comments != null) ? this.photos.comments.length : 0
        });
    },
    components: { CommentModal }
}
</script>

<template>

    

	<div class="container">
        <!--This modal is only visible when we open the comments-->
        <CommentModal v-if="isModalVisible" :comms="this.photos.comments" :profileId = "this.photos.profileId" :photoId = "this.photos.photoId" :id = "this.myId" @close="closeModal"/>
        
        <!--This row contains a small photo optios menu and photo profile name-->
        <div class="row ">
            <div class="col">
                <div style="float:left;">
                    <button class="btn" @click="goToProfile" type="button" style="padding: 0; border:0;">
                        <svg class="feather" style="stroke-width:2">
                            <use href="/feather-sprite-v4.29.0.svg#user"/>
                        </svg>        
                    </button>
                    {{this.photos.profileId}}
                    
                </div>
            </div>
            <div class="col">
                <!-- Default dropright button -->
                <div class="dropdown" style="float:right;">
                    <button class="btn" type="button" id="dropdownMenuButton1" data-bs-toggle="dropdown" aria-expanded="false" style="padding: 0; border:0;">
                        <svg class="feather" style="stroke-width:2">
                            <use href="/feather-sprite-v4.29.0.svg#more-vertical" />
                        </svg>
                    </button>
                    <ul class="dropdown-menu" aria-labelledby="dropdownMenuButton1">
                        <li v-if="this.photos.profileId != this.myId">
                            <a class="dropdown-item" @click="blockUser">Block User</a>
                        </li>
                        <li v-else>
                            <a class="dropdown-item" @click="deletePhoto">Delete Photo</a>
                        </li>
                    </ul>
                </div>
            </div>
            
        </div>
        <!--First row contains the image-->       
        <div class="row justify-content-center">
            <img v-bind:src = "'data:image/gif;base64,' + this.photos.image ">
        </div>
        <!--Second row contains likes and comment-->
        <div class="row justify-content-around">
            <div class="col text-center">
                <li class="nav-item" >
                    <label> {{this.likeCount}}</label>
                    <!--The v-if checks if the user that is visualizing the image have already left a like to it-->
                    <svg class="feather" v-if="this.likeFlag" @click="unlikePhoto"><use href="/feather-sprite-v4.29.0.svg#heart" style="fill:black"/></svg> 
                    <!--The else means that you DID NOT leave a like yet-->
                    <svg class="feather" v-else @click="likePhoto"><use href="/feather-sprite-v4.29.0.svg#heart"/></svg> 

                </li>
            </div>
            <div class="col text-center">
                <li class="nav-item">
                    <label>{{this.commentsCount}}</label>
                    <svg class="feather" @click="openCommentModal"><use href="/feather-sprite-v4.29.0.svg#message-square"/></svg> 
                </li>
            </div>                   
        </div>
    </div>
</template>

<style></style>