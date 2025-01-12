<script lang="ts">
const { appName } = $props();
</script>

<header class="hero is-info">
    <div class="hero-body has-text-centered">
        <p class="title is-capitalized">{appName}</p>
    </div>
</header>

<style>
    p {
        color: white;
    }
</style>
