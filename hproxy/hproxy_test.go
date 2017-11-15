func handleGetPets(w http.ResponseWriter, r *http.Request) {
	url, _ := url.Parse(animalServiceBasePath)
	proxy := goengine.NewSingleHostReverseProxy(url)

	r.Header.Add("x-auth", servicesAuthorizationKey)
	r.URL.Path = fmt.Sprintf("/users/%s/animals", r.Context().Value("userID").(string))

	proxy.ServeHTTP(w, r)
}