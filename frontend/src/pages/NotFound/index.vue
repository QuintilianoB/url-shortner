<template>

	<b-container fluid>
	
		<b-row>
			<b-col class="container-flex" style="height: 92vh">
					
                <b-col>
                    
                    <b-col offset-md="2" md="8" class="p-2">
                        <img style="width: 95%" src="../../assets/bg-404.svg" alt="">
                    </b-col>

                </b-col>                
            </b-col>
			<b-col>
				<b-col class="container-flex" style="height: 92vh">
					
					<b-col>						
                        <b-col md="12" class="p-2">
                            <p class="title01">Ops!</p>
                            <p class="title02">Página não encontrada.</p>
                            <p class="title03">
                                Alguma coisa aconteceu, mas não se preocupe! <br>Clique no botão abaixo para voltar a página inicial.
                            </p>
                        </b-col>
						
                        <b-col md="12" class="p-2 title-login-portal" style="margin-top: 70px; text-align: center;">
                            <router-link to="/">
                                <b-button class="button-notFound">Página Inicial</b-button>
                            </router-link>
                        </b-col>					

					</b-col>
				</b-col>
			</b-col>

		</b-row>
        <b-row class="row-copyright">
            <b-col>
                <span>Copyright ® Soluti Certificação Digital</span>
            </b-col>
        </b-row>

	</b-container>
</template>

<script>
export default {
    // created: function() {
    //     // Redirect outside the app using plain old javascript
    //     window.location.href = "/my-new-404-page.html";
    // }
}
</script>

<style>
@import "./styles.css";
</style>
