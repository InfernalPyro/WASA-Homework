<script setup>
import UserItem from '../components/UserItem.vue';

</script>
<script>


export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
            users: null,
            flag : false,
		}
	},
	methods: {
        async search() {
            this.flag = false;
			this.loading = true;
			this.errormsg = null;
			this.userToFind = document.getElementById('searchBox').value
			try {
				let response = await this.$axios({
                    method: "get",
                    url: "/search/" + this.userToFind, 
                    headers: {
                        "Content-Type": "application/json",
                        "Authorization": "Bearer " + this.token,
                    }
                });
				this.users = response.data;
                this.flag = true;
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
		
	},
	mounted() {
        this.$nextTick(function(){
            
        })		
	},
	components: { UserItem }


}
</script>

<template>
	<!--This column contains all the photos in the stream--> 
    <div class="col-8 text-center" style="padding-left: 12%;">
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<div class="row">
			Find whoever you're looking for by typing their name here:
		</div>
		<br><br>
		
		<div class="row">
			<input type="text" id="searchBox"  placeholder="Search.." @input="search">
		</div>

        <div id="rbHiderFillTop"></div>
        <!--Each container is made of the photo and the likes and comment buttons-->	    
        <div>
            <UserItem v-if="flag" v-for= "user in users" :id = "user.id" :name = "user.username" :key="user.UserId"></UserItem>
        </div>
	</div>

</template>

<style>
</style>