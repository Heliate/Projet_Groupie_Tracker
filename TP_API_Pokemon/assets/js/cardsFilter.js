// Attendre que le DOM soit complètement chargé
document.addEventListener('DOMContentLoaded', function() {
    // Sélectionner les éléments du DOM
    const applyFiltersButton = document.getElementById('applyFiltersButton');
    const resetFiltersButton = document.getElementById('resetFiltersButton');

    // Appliquer les filtres et mettre à jour l'URL
    applyFiltersButton.addEventListener('click', () => {
        const formData = new FormData(document.getElementById('filterForm'));
        const rarityFilter = formData.getAll('rarity_filter');
        const typeFilter = formData.getAll('type_filter');
        const weaknessFilter = formData.getAll('weakness_filter');
        const sortFilter = formData.get('sort_filter');

        const url = new URL(window.location);
        
        // Préserver le paramètre set et page
        const setParam = url.searchParams.get('set');
        const pageParam = url.searchParams.get('page');
        const searchParam = url.searchParams.get('search');
        
        // Supprimer tous les paramètres
        url.search = '';
        
        // Restaurer les paramètres importants
        if (setParam) {
            url.searchParams.set('set', setParam);
        }
        if (pageParam) {
            url.searchParams.set('page', pageParam);
        }
        if (searchParam) {
            url.searchParams.set('search', searchParam);
        }
        
        // Ajouter les filtres
        rarityFilter.forEach(rarity_filter => {
            url.searchParams.append('rarity_filter', rarity_filter);
        });
        typeFilter.forEach(type_filter => {
            url.searchParams.append('type_filter', type_filter);
        });
        weaknessFilter.forEach(weakness_filter => {
            url.searchParams.append('weakness_filter', weakness_filter);
        });
        if (sortFilter) {
            url.searchParams.append('sort_filter', sortFilter);
        }
        
        // Rediriger vers la nouvelle URL avec les filtres
        window.location.href = url.toString();
    });

    // Réinitialiser les filtres
    resetFiltersButton.addEventListener('click', () => {
        const form = document.getElementById('filterForm');
        form.reset();

        const url = new URL(window.location);
        
        // Préserver le paramètre set et page
        const setParam = url.searchParams.get('set');
        const pageParam = url.searchParams.get('page');
        const searchParam = url.searchParams.get('search');
        
        // Supprimer tous les paramètres
        url.search = '';
        
        // Restaurer les paramètres importants
        if (setParam) {
            url.searchParams.set('set', setParam);
        }
        if (pageParam) {
            url.searchParams.set('page', pageParam);
        }
        if (searchParam) {
            url.searchParams.set('search', searchParam);
        }
        
        // Rediriger vers la nouvelle URL sans filtres
        window.location.href = url.toString();
    });
});