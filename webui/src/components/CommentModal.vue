<script>
import CommentItem from './CommentItem.vue';


export default {
    name: 'modal',
    props: ['comms', 'profileId', 'photoId', 'id'],
    data: function() {
		return {
            prId: this.profileId,
            phId: this.photoId,
            comments: this.comms,
            myId : this.id,
            token: sessionStorage.getItem("storedData"), 
        }
    },
    methods: {
        close() {
            this.$emit('close');
        },
        async commentPhoto() {
            this.loading = true;
            this.errormsg = null;
            try {

                let config ={
                    headers:{
                        'Content-Type': "application/json",
                        'Authorization': "Bearer " + this.token,

                    }
                }
                let data ={
                    "userId" : Number(this.myId),
                    "comment": document.getElementById('newComment').value
                }
                let response = await this.$axios.post(
                    "/profile/" + this.prId + "/photo/" + this.phId + "/comments/", data, config);
              
            }
            catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        }

    },
    components: {CommentItem }
};
</script>
<template>
    <transition name="modal-fade">
        <div class="modal-backdrop">
            <div
                class="modal"
                role="dialog"
                aria-labelledby="modalTitle"
                aria-describedby="modalDescription"
            >
                <header class="modal-header" id="modalTitle">
                    Comments
                    <button type="button" class="btn-close" @click="close" aria-label="Close Modal App">
                        <svg class="feather" style="stroke:red; stroke-width:5"><use href="/feather-sprite-v4.29.0.svg#x"/></svg>
                    </button>
                </header>
                <!--Body of the modal (Contain every comment of this photo)-->
                <section class="modal-body" id="modalDescription">
                    <div>
                        <CommentItem v-for= "comment in comments" :comm="comment" :userId="this.myId" :profileId="this.prId" ></CommentItem>
                    </div>
                </section>
                <!--Make a new comment box-->
                <footer class="modal-footer">
                    <input type="text" id="newComment" placeholder="Write a new comment">
                    <button @click="commentPhoto()">Publish</button>
                </footer>
                
            </div>
        </div>
    </transition>
</template>


<style>


.modal-backdrop {
  position: fixed; /* Stay in place */
  left: 0;
  top: 0;
  width: 100%; /* Full width */
  height: 100%; /* Full height */
  overflow: auto; /* Enable scroll if needed */
  background-color: rgb(0,0,0); /* Fallback color */
  background-color: rgba(0,0,0,0.4); /* Black w/ opacity */
}
.modal {
    background: #ffeeee;
    box-shadow: 4px 4px 30px 2px;
    overflow-x: auto;
    display: flex;
    flex-direction: column;
    height: 30vw;
    width: fit-content;
    top: 10%;
    left: 30%;
    right: auto;
    overflow-y: scroll; 
}
.modal-header,
.modal-footer {
    padding: 10px;
    display: flex;
    
}
.modal-header {
    border-bottom: 1px solid #eeeeee;
    color: #4aae9b;
    justify-content: space-between;
}
.modal-footer {
    border-top: 2px solid #eeeeee;
    justify-content: flex-end;
    position: sticky;
    bottom: 0;
    background-color: #eeeeee ;
}
.modal-body {
    position: relative;
    padding: 30px 20px;
}
.btn-close {
    border: 2px solid #ae4a57;
    font-size: 25px;
    padding: 25px;
    cursor: pointer;
    font-weight: bold;
    color: #4aae9b;
    background: transparent;
}


</style>
