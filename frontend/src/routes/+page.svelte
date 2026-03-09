<script lang="ts">
    import { default as ArtGrid } from "$lib/components/artGrid/ArtGrid.svelte"

</script>

<div class="art-page">
    <h1>Welcome to Portfolio</h1>
    <div>
        <ArtGrid />
    </div>
</div>

<style>
    .art-page {
        font-family: 'CustomFont', sans-serif;
        width: full;
    }
</style>
