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
    let radius = 400; 
    let animating = false;
    
    // Créer les éléments du carrousel
    function createCarouselItems() {
        // Vider le carousel d'abord
        carousel.innerHTML = '';
        
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
        });
        
        // Positionner tous les éléments
        updateCarousel();
        
        // Afficher le bouton d'ouverture
        if (openBtn) openBtn.style.display = 'block';
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
            
            // Ajuster l'opacité en fonction de la distance
            const opacity = Math.max(0.3, 1 - Math.abs(angle) / 180);
            item.style.opacity = opacity;
            
            // Ajuster le z-index pour que les éléments au premier plan soient au-dessus
            const zIndex = Math.round(100 - Math.abs(angle));
            item.style.zIndex = zIndex;
            
            // Vérifier si c'est l'élément sélectionné (à l'avant)
            if (index === currentIndex) {
                item.classList.add('selected');
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
    if (prevBtn) prevBtn.addEventListener('click', rotatePrev);
    if (nextBtn) nextBtn.addEventListener('click', rotateNext);
    if (openBtn) openBtn.addEventListener('click', openBooster);
    
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
    
    // Support tactile pour mobile
    let touchStartX = 0;
    let touchEndX = 0;
    
    carousel.addEventListener('touchstart', function(e) {
        touchStartX = e.changedTouches[0].screenX;
    });
    
    carousel.addEventListener('touchend', function(e) {
        touchEndX = e.changedTouches[0].screenX;
        handleSwipe();
    });
    
    function handleSwipe() {
        const minSwipeDistance = 50;
        if (touchEndX < touchStartX - minSwipeDistance) {
            // Swipe vers la gauche
            rotateNext();
        } else if (touchEndX > touchStartX + minSwipeDistance) {
            // Swipe vers la droite
            rotatePrev();
        }
    }
    
    // Initialiser le carrousel
    createCarouselItems();
    
    // Ajuster le rayon selon la taille de l'écran
    function adjustRadius() {
        const containerWidth = document.querySelector('.carousel').clientWidth;
        radius = Math.min(400, containerWidth * 0.45);
        updateCarousel();
    }
    
    // Appeler au redimensionnement
    window.addEventListener('resize', adjustRadius);
    
    // Appeler une fois au chargement
    adjustRadius();
});