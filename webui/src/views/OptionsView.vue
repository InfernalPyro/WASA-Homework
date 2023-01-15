<script>

export default {

    data: function () {
        return {
            errormsg: null,
            loading: false,
            token : sessionStorage.getItem('storedData'),
            userId: this.$route.params.userId,
        };
    },
    methods: {
        async changeUsername() {
            this.loading = true;
            this.errormsg = null;
            this.username = document.getElementById('newUsername').value
            
            if(this.username.length < 3 || this.username.length > 16){
				alert("Username must have more than 3 chars and less than 16")
				return
			}

            try {
                let config ={
                    headers:{
                        'Content-Type': "application/json",
                        "Authorization": "Bearer " + this.token,

                    }
                }
                let data ={
                    "username" : this.username
                }

				let response = await this.$axios.patch("/user/" + this.userId + "/username", data, config);
                this.$router.push("/");
            }
            catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },
    },
    mounted() {
        this.$nextTick(function () {

        });
    },

}


</script>

<template>
    <div class="row">
        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
        <div class="col-auto">
            <label>Type here to change your username : </label>
        </div>
        <div class="col-auto">
            <input type="text" id="newUsername"  minlenght = "3" maxlenght = "16" pattern = "^[a-zA-Z0-9 ]*$" style="margin-right: 4vw;">
	        <button @click="changeUsername">Confirm</button>
        </div>
    </div>
</template>

<style>
</style>
