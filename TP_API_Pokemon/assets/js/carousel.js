document.addEventListener('DOMContentLoaded', function() {
    // Vérifier si le carrousel existe sur cette page
    const carousel = document.getElementById('boosterCarousel');
    if (!carousel) return;
    
    const prevBtn = document.getElementById('prevBtn');
    const nextBtn = document.getElementById('nextBtn');
    const openBtn = document.getElementById('openBooster');
    
    // Configuration des boosters avec leurs liens correspondants
    const boosterTypes = [
        { id: 1, name: "Sword & Shield", image: "/static/images/swsh1.png", link: "/cards?set=swsh1" },
        { id: 2, name: "Rebel Clash", image: "/static/images/swsh2.png", link: "/cards?set=swsh2" },
        { id: 3, name: "Darkness Ablaze", image: "/static/images/swsh3.png", link: "/cards?set=swsh3" },
        { id: 4, name: "Vivid Voltage", image: "/static/images/swsh4.png", link: "/cards?set=swsh4" },
        { id: 5, name: "Battle Styles", image: "/static/images/swsh5.png", link: "/cards?set=swsh5" },
        { id: 6, name: "Chilling Reign", image: "/static/images/swsh6.png", link: "/cards?set=swsh6" },
        { id: 7, name: "Evolving Skies", image: "/static/images/swsh7.png", link: "/cards?set=swsh7" },
        { id: 8, name: "Fusion Strike", image: "/static/images/swsh8.png", link: "/cards?set=swsh8" },
    ];
    
    let currentIndex = 0;
    let radius = 300; // Rayon du cercle
    let animating = false;
    
    // Créer les éléments du carrousel
    function createCarouselItems() {
        boosterTypes.forEach((booster, index) => {
            const item = document.createElement('div');
            item.className = 'carousel-item';
            item.dataset.id = booster.id;
            
            // Ajouter l'image
            const img = document.createElement('img');
            img.src = booster.image;
            img.alt = booster.name;
            
            item.appendChild(img);
            carousel.appendChild(item);
            
            // Ajouter un gestionnaire d'événements de clic
            item.addEventListener('click', function() {
                if (!animating) {
                    const itemId = parseInt(this.dataset.id);
                    const targetIndex = boosterTypes.findIndex(b => b.id === itemId);
                    rotateToIndex(targetIndex);
                }
            });
        });
        
        // Positionner tous les éléments
        updateCarousel();
    }
    
    // Mettre à jour la position des éléments du carrousel
    function updateCarousel() {
        const items = document.querySelectorAll('.carousel-item');
        const angleStep = 360 / items.length;
        
        items.forEach((item, index) => {
            // Calculer la position (angle) pour chaque élément
            const angle = angleStep * ((index - currentIndex + items.length) % items.length);
            const radians = (angle * Math.PI) / 180;
            
            // Calculer la position sur le cercle
            const x = radius * Math.sin(radians);
            const z = radius * Math.cos(radians) - radius;
            
            // Appliquer la transformation
            item.style.transform = `translate(calc(-50% + ${x}px), -50%) translateZ(${z}px)`;
            
            // Vérifier si c'est l'élément sélectionné (à l'avant)
            if (index === currentIndex) {
                item.classList.add('selected');
                openBtn.style.display = 'block';
            } else {
                item.classList.remove('selected');
            }
        });
    }
    
    // Rotation vers un index spécifique
    function rotateToIndex(index) {
        if (animating || index === currentIndex) return;
        
        animating = true;
        currentIndex = index;
        updateCarousel();
        
        // Rétablir l'état après l'animation
        setTimeout(() => {
            animating = false;
        }, 500); // Correspondant à la durée de transition CSS
    }
    
    // Rotation vers l'élément suivant
    function rotateNext() {
        if (animating) return;
        const nextIndex = (currentIndex + 1) % boosterTypes.length;
        rotateToIndex(nextIndex);
    }
    
    // Rotation vers l'élément précédent
    function rotatePrev() {
        if (animating) return;
        const prevIndex = (currentIndex - 1 + boosterTypes.length) % boosterTypes.length;
        rotateToIndex(prevIndex);
    }
    
    // Fonction d'ouverture du booster - redirection vers la page correspondante
    function openBooster() {
        if (animating) return;
        
        const selectedItem = document.querySelector('.carousel-item.selected');
        if (selectedItem) {
            animating = true;
            selectedItem.classList.add('opening-animation');
            
            setTimeout(() => {
                selectedItem.classList.remove('opening-animation');
                animating = false;
                
                // Récupérer l'ID du booster sélectionné
                const boosterId = parseInt(selectedItem.dataset.id);
                const booster = boosterTypes.find(b => b.id === boosterId);
                
                // Rediriger vers la page correspondante
                if (booster && booster.link) {
                    window.location.href = booster.link;
                }
                
            }, 1000);
        }
    }
    
    // Gestionnaires d'événements
    prevBtn.addEventListener('click', rotatePrev);
    nextBtn.addEventListener('click', rotateNext);
    openBtn.addEventListener('click', openBooster);
    
    // Support clavier
    document.addEventListener('keydown', function(e) {
        if (e.key === 'ArrowLeft') {
            rotatePrev();
        } else if (e.key === 'ArrowRight') {
            rotateNext();
        } else if (e.key === 'Enter') {
            openBooster();
        }
    });
    
    // Initialiser le carrousel
    createCarouselItems();
});