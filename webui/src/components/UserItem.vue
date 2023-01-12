
<script>


export default {
    props: ['id','name'],
    data: function() {
		return {
            userId: this.id,
            username : this.name,
        }
    },
    methods: {
        async goToProfile(){
            this.$router.push("/user/" + this.userId + "/profile");
        },

    },

    created() {

    },
};
</script>
<template>
    <div class="container">
        <button class="btn" @click="goToProfile" type="button" style="padding: 0; border:0;">
            <div class="row justify-content-start">
                <div class="col-auto">
                    <svg class="feather" style="stroke-width:2">
                        <use href="/feather-sprite-v4.29.0.svg#user"/>
                    </svg>        
                    {{this.username}}
                </div>
            </div>
        </button>
    </div>

</template>