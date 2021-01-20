<template>
	<view>
		<form action="">
			<input type="text" placeholder="请输入手机号" v-model="user.telephone">
			<text :v-if="!msg">{{msg}}</text>
			<input type="password" placeholder="请输入密码" v-model="user.password">
		</form>
		<button type="default" @click="login">登录</button>
	</view>
</template>

<script>
	// 引入md5js插件
	import md5 from '../../static/js/md5.js'
	export default {
		data() {
			return {
				user: {
					telephone: '',
					password: ''
				},
				msg: ''
			};
		},
		methods: {
			login() {
				const user = this.user
				uni.request({
					header: {
						'Content-Type': 'application/x-www-form-urlencoded'
					},
					url: this.$api + "/auth/login",
					method: 'POST',
					data: {
						telephone: user.telephone,
						password: md5(user.password)
					},
					dataType: 'json',
					success: (res) => {
						console.log(res.data)
						if (res.data.code === 200) {
							uni.setStorageSync('token', res.data.data.token)
							uni.showToast({
								icon: 'none',
								title: '登录成功',
								// #ifdef MP-WEIXIN
								duration: 1000,
								// #endif
								// mask: true
							});
							setTimeout(function() {
								uni.reLaunch({
									url: '/pages/home/home'
								});
							}, 1000);
						}else{
							this.msg = res.data.msg
						}
					}
				})
			}
		}
	}
</script>

<style lang="scss" scoped>
	view {
		width: 80%;
		position: absolute;
		top: 15%;
		left: 50%;
		transform: translate(-50%, 0);

		form {
			input {
				border: 1px solid #DADCE0;
				margin: 25rpx auto;
			}
		}

		button {
			margin-top: 50rpx;
			font-size: 28rpx;
			border-radius: 50rpx;
			border: 1rpx solid #DADCE0;
		}
	}
</style>

