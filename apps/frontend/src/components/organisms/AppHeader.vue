<script setup lang="ts">
import { useRouter } from "vue-router";
import { useAuthStore } from "../../store/auth.store";
import { message } from "ant-design-vue";

const auth = useAuthStore();
const router = useRouter();

const handleLogout = () => {
	auth.logout();
	message.success("Logged out successfully");
	router.push("/login");
};
</script>

<template>
	<a-layout-header class="wrapper-header">
		<div class="logo">Payment Dashboard</div>
		<a-button
			class="btn"
			@click="handleLogout"
			>Logout</a-button
		>
	</a-layout-header>
</template>

<style scoped>
.wrapper-header {
	display: flex;
	justify-content: space-between;
	align-items: center;
	background: #00949e;
}

.logo {
	color: white;
	font-weight: bold;
	font-size: 18px;
}

.btn {
	background: #fbc037;
	border: none;
	color: black;
	font-weight: 600;
}

.btn:hover {
	background: #f9a825;
	color: black;
}
</style>
