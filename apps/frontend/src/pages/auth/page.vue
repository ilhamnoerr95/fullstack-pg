<script setup lang="ts">
import LoginForm from "../../components/molecules/LoginForm.vue";
</script>

<template>
	<div class="login-page">
		<LoginForm />
	</div>
</template>

<style scoped>
.login-page {
	height: 90vh;
	display: flex;
	flex-direction: row;
	justify-content: center;
	align-items: center;
	min-width: 320px;
}
</style>
