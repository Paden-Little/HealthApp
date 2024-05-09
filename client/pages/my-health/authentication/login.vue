<script setup lang="ts"></script>
<template>
	<div class="w-96 flex justify-center flex-col content-center">
		<h1 class="text-2xl font-bold">My Health Login</h1>
		<form
			class="flex flex-col justify-center content-center"
			action=""
			method="post"
		>
			<label for="email">Email</label>
			<input type="email" name="email" id="email" />
			<label for="password">Password</label>
			<input type="password" name="password" id="password" />
			<button type="submit">Login</button>
		</form>
	</div>
</template>
