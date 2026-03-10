<script lang="ts">
    import { default as ArtTile } from "$lib/components/artTile/ArtTile.svelte"
    import type { ArtTileInter } from "../artTile/ArtTileInterface";

    let tiles: ArtTileInter[] = $state([
        { Id: 1, Title: 'Art Title 1', URL: 'Info', Description: "", Portrait: false},
        { Id: 2, Title: 'Art Title 2', URL: 'Info', Description: "", Portrait: false},
        { Id: 3, Title: 'Art Title 3', URL: 'Info', Description: "", Portrait: false},
        { Id: 4, Title: 'Art Title 4', URL: 'Info', Description: "", Portrait: false},
    ]);

    
    $effect(() => {
        (async () => {
            const res = await fetch(`http://localhost:8080/api/art/`);
            tiles = await res.json();
        })();
    });
</script>

<div class="grid-area">
    {#each tiles as tile (tile.Id)}
        <ArtTile title={tile.Title} url={tile.URL} portrait={tile.Portrait} />
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

    @media screen and (max-width: 450px) {
        .grid-area {
            display: grid;
            grid-template-columns: repeat(1, 1fr);
        }
    }
</style>
