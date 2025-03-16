// Attendre que le DOM soit complètement chargé avant d'exécuter le code
document.addEventListener('DOMContentLoaded', function() {
    // Sélectionner les boutons de filtrage du DOM
    const applyFiltersButton = document.getElementById('applyFiltersButton');
    const resetFiltersButton = document.getElementById('resetFiltersButton');

    // Gestionnaire d'événement pour le bouton d'application des filtres
    applyFiltersButton.addEventListener('click', () => {
        // Récupération des données du formulaire de filtrage
        const formData = new FormData(document.getElementById('filterForm'));
        const rarityFilter = formData.getAll('rarity_filter');
        const typeFilter = formData.getAll('type_filter');
        const weaknessFilter = formData.getAll('weakness_filter');
        const sortFilter = formData.get('sort_filter');

        // Création d'un objet URL pour manipuler les paramètres d'URL
        const url = new URL(window.location);
        
        // Préserver les paramètres importants existants (set, page, recherche)
        const setParam = url.searchParams.get('set');
        const pageParam = url.searchParams.get('page');
        const searchParam = url.searchParams.get('search');
        
        // Supprimer tous les paramètres actuels pour repartir proprement
        url.search = '';
        
        // Restaurer les paramètres importants qui doivent être conservés
        if (setParam) {
            url.searchParams.set('set', setParam);
        }
        if (pageParam) {
            url.searchParams.set('page', pageParam);
        }
        if (searchParam) {
            url.searchParams.set('search', searchParam);
        }
        
        // Ajouter les nouveaux filtres sélectionnés à l'URL
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
        
        // Rediriger vers la nouvelle URL avec les filtres appliqués
        window.location.href = url.toString();
    });

    // Gestionnaire d'événement pour le bouton de réinitialisation des filtres
    resetFiltersButton.addEventListener('click', () => {
        // Réinitialiser tous les champs du formulaire
        const form = document.getElementById('filterForm');
        form.reset();

        // Création d'un objet URL pour manipuler les paramètres d'URL
        const url = new URL(window.location);
        
        // Préserver les paramètres importants existants (set, page, recherche)
        const setParam = url.searchParams.get('set');
        const pageParam = url.searchParams.get('page');
        const searchParam = url.searchParams.get('search');
        
        // Supprimer tous les paramètres de filtre actuels
        url.search = '';
        
        // Restaurer uniquement les paramètres importants
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