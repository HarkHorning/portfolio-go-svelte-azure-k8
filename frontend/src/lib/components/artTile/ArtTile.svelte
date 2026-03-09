<script lang="ts">
    // const title = "Artwork Name"
    // export const title = "Artwork Name";
    let { title } : { title: string } = $props()
</script>

<div class="card">
  <h2>{title}</h2>
</div>

<style>
  .card { border: 1px solid #ccc; padding: 1rem; border-radius: 8px; }
</style>