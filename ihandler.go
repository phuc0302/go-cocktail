package cocktail

/**
 * IHandler represent a routing handler function.
 *
 *  Injector will inject these parameters dynamically as function's inputs
 *    + http.Header           (Optional)
 *    + *http.Request         (Optional)
 *    + http.ResponseWriter   (Optional)
 *
 *    + url.Values            (Optional)
 *    + cocktail.FileParams   (Optional)
 *    + cocktail.PathParams   (Optional)
 *
 *
 *  Function should only return one or two parameter(s)
 *    + cocktail.HttpStatus   (Optional)
 *    + struct or string      (Optional)
 *    + template              (Optional)  (html/template)
 */
type IHandler interface{}
