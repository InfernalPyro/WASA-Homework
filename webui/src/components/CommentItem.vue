
<script>


export default {
    props: ['comm','userId','profileId'],
    data: function() {
		return {
            errormsg: null,
            comment: this.comm,
            myId: this.userId,
            prId: this.profileId,
            token: sessionStorage.getItem("storedData"), 

        }
    },
    methods: {
        async blockUser() {
            this.loading = true;
            this.errormsg = null;
            try {
                let response = await this.$axios({
                    method: "put",
                    url: "/user/" + this.myId + "/ban/" + this.comment.userId, 
                    headers: {
                        "Content-Type": "application/json",
                        "Authorization": "Bearer " + this.token,
                    }
                });
            }
            catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },
        async deleteComment() {
            this.loading = true;
            this.errormsg = null;
            try {
                let response = await this.$axios({
                    method: "delete",
                    url: "/profile/" + this.prId + "/photo/" + this.comment.photoId + "/comments/" + this.comment.commentId,
                    headers: {
                        "Content-Type": "application/json",
                        "Authorization": "Bearer " + this.token,
                    }
                });
                this.$emit('deleted');
            }
            catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },
    },

    created() {

    },
};
</script>
<template>
    <div class="container" style="width: 30vw;">
        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
        <div class="row justify-content-end">
            <div class="col text-start">{{this.comment.username}}</div>
            <div class="col-2">
                <!-- Default dropright button -->
                <div class="dropdown">
                    <button class="btn" @click="options" type="button" id="dropdownMenuButton1" data-bs-toggle="dropdown" aria-expanded="false" style="padding: 0; border:0;">
                        <svg class="feather" style="stroke-width:2">
                            <use href="/feather-sprite-v4.29.0.svg#more-vertical" />
                        </svg>
                    </button>
                    <ul class="dropdown-menu" aria-labelledby="dropdownMenuButton1">
                        <template v-if="this.comment.userId != this.myId">
                            <li >
                                <a class="dropdown-item" @click="blockUser">Block User</a>
                            </li>
                        </template>
                        <template v-else>
                            <li>
                                <a class="dropdown-item" @click="deleteComment">Delete Comment</a>
                            </li>
                        </template>   
                    </ul>
                </div>
            </div>
        </div>
        <div id="rbHiderFillTop" style="height: 20px;"></div>

        {{ this.comment.comment }}
        <footer>{{ this.comment.time }}</footer>
    </div>
</template>