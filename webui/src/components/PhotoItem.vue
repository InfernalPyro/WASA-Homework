<script>
export default {
    props: ['images'],
    data() {
        return {
            photos: this.images,
        }
    },
    created() {
    },
}
</script>

<template>
	    <div class="container">
            <!--First row contains the image-->
            <div class="row justify-content-center">
                <img v-bind:src = "'data:image/gif;base64,' + this.photos.image ">
            </div>
            <!--Second row contains likes and comment-->
            <div class="row justify-content-around">
                <div class="col text-center">
                    <li class="nav-item">
                        <label>{{(this.photos.likes != null) ? this.photos.likes.length : 0}}</label>
                        <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#heart"/></svg> 
                    </li>
                </div>
                <div class="col text-center">
                    <li class="nav-item">
                        <label>{{(this.photos.comments != null) ? this.photos.comments.length : 0}}</label>
                        <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#message-square"/></svg> 
                    </li>
                </div>                   
            </div>
        </div>
</template>

<style></style>