<script lang="ts">
    import { default as ArtTile } from "$lib/components/artTile/ArtTile.svelte"
    import type { ArtTileInter } from "../artTile/ArtTileInterface";

    const tiles: ArtTileInter[] = [
        { id: 1, title: 'Art Title 1', url: 'Info' },
        { id: 2, title: 'Art Title 2', url: 'Info' },
        { id: 3, title: 'Art Title 3', url: 'Info' },
    ];
</script>

<div class="grid-area">
    {#each tiles as tile (tile.id)}
        <ArtTile title={tile.title}/>
    {/each}
</div>

<style>
    .grid-area {
        display: grid;
        grid-template-columns: repeat(3, 1fr);
    }

    @media screen and (max-width: 850px) {
        .grid-area {
            display: grid;
            grid-template-columns: repeat(2, 1fr);
        }
    }
</style>
