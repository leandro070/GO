$(function () {
    var espacioNombre = {};


    (function (app) {

        app.init = function () {
            app.limpiarFormulario();
            app.obtenerProvincias();
            app.bindings();

        };

        app.bindings = function () {
            document.getElementById("provincia").addEventListener("change", app.obtenerProvincias());

            // oyente del boton guardar
            $("#guardar").on("click", function (event) {
                app.validarRegistro();
            });
            
            $('#calendario_nacimiento').datepicker({
                    format: "yyyymmdd"
            });
        };
      
            
        
        
        app.obtenerProvincias = function () {
            document.getElementById("provincia");

            var url = "controlador/Ruteador.php?accion=Provincia&id=-1";

            $.ajax({
                type: "POST",
                dataType:'json',
                url: url,
                success: function (datosRecibidos)
                {    
                    app.actualizarProvincia(datosRecibidos);
                },
                error: function(datosRecibidos){
                    alert('ERROR: provincia');
                    alert(datosRecibidos);
                }
            });
        };
        
        $('#provincia').on('change', function() {
            var OID_provincia= this.value;
            var url = "controlador/Ruteador.php?accion=Departamento&id="+OID_provincia;
            $.ajax({
                type: "POST",
                dataType:'json',
                url: url,

                success: function (datosRecibidos)
                {    
                    app.actualizarDepartamento(datosRecibidos);
                },
                error: function(datosRecibidos){
                    alert('ERROR: departamento');
                    alert(datosRecibidos);
                }
            });

        })
         
        app.limpiarFormulario = function () {    //funcion para limpiar los textbox del modal
            $("#nombre").val('');
            $("#apellido").val('');
            $("#apellidomaterno").val('');
            $("#dni").val(0);
            $("#cuil").val('');
            $("#FechaNacimiento_nacimiento").val('');
            $("#correo").val('');
            $("#telefono").val('');
            $("#telefonoa").val('');
            $("#celular").val('');
            $("#direccion").val('');
            $("#numero").val(0);
            $("#piso").val(0);
            $("#dpto").val(0);
            
        };

        app.validarRegistro = function(){
            $('#form-persona').validator();
            $('#form-domicilio').validator();
            var listoParaEnviar=true;
            $('#form-persona').validator('validate');
            $('#form-domicilio').validator('validate');


            e = document.getElementById("departamento");
            var departamento = e.options[e.selectedIndex].value;
            p = document.getElementById("provincia");
            var provincia = p.options[p.selectedIndex].value;
            var nombre=document.getElementById('nombre').value;
            var apellido=document.getElementById('apellido').value;
            var dni=document.getElementById('dni').value;
            var cuil=document.getElementById('cuil').value;
            var fnacimiento=document.getElementById('fecha_nacimiento').value;
            var email=document.getElementById('correo').value;
            var direccion=document.getElementById('direccion').value;
            var numero=document.getElementById('numero').value;
            var depto=document.getElementById('depto').value;
            var apellido_materno=document.getElementById('apellidomaterno').value;
            var telefonofijo=document.getElementById('telefono').value;
            var celular=document.getElementById('celular').value;
            var telefonoalter=document.getElementById('telefonoa').value;
            var piso = document.getElementById('piso').value;
            var dpto = document.getElementById('depto').value;


            var caracteres_numeros="0123456789";
            var caracteres_letras="abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";
            var carateres_especiales="!·$%&/()=?¿*^¨Ç,.-{}[]~¬½~#|ºª><";
            var arroba="@";


            //Validar nombre
            if (nombre.length==0 || nombre.length>45) {
                listoParaEnviar=false;

            }   
            var input=nombre;
            for (var i = input.length - 1; i >= 0; i--) {
                var numero_ilegal=false;
                var caracter_ilegal=false;
                for (var a = carateres_especiales.length - 1; a >= 0; a--) {
                    if (input[i]==carateres_especiales[a] || input[i]==arroba) {
                        listoParaEnviar=false;
                        S
                        caracter_ilegal=true;
                        break;                        
                    }                    

                }
                for (var o = caracteres_numeros.length - 1; o >= 0; o--) {
                    if (input[i]==caracteres_numeros[o]){
                        listoParaEnviar=false;
                        numero_ilegal=true;
                        break;
                    }
                }

                if (numero_ilegal) {
                    alert("El campo NOMBRE tiene caracteres NUMERICOS ilegales, asegurate de usar solo letras ");
                    break;
                }
                if (caracter_ilegal) {
                    alert("El campo NOMBRE tiene caracteres ESPECIALES ilegales, asegurate de usar solo letras ");
                    break;
                }

            }


            //Validar apellido
            if (apellido.length==0 || apellido.length>45) {

                listoParaEnviar=false;
            }

            var input=apellido;
            for (var i = input.length - 1; i >= 0; i--) {
                var numero_ilegal=false;
                var caracter_ilegal=false;
                for (var a = carateres_especiales.length - 1; a >= 0; a--) {
                    if (input[i]==carateres_especiales[a] || input[i]==arroba) {
                        listoParaEnviar=false;
                        caracter_ilegal=true;
                        break;                        
                    }                    

                }
                for (var o = caracteres_numeros.length - 1; o >= 0; o--) {
                    if (input[i]==caracteres_numeros[o]){
                        listoParaEnviar=false;
                        numero_ilegal=true;
                        break;
                    }
                }

                if (numero_ilegal) {
                    alert("El campo APELLIDO tiene caracteres NUMERICOS ilegales, asegurate de usar solo letras ");
                    break;
                }
                if (caracter_ilegal) {
                    alert("El campo APELLIDO tiene caracteres ESPECIALES ilegales, asegurate de usar solo letras ");
                    break;
                }

            }

            if(apellido_materno.length>45){
                listoParaEnviar=false;
            }
            //Validar dni
            if (dni.length==0||dni.length>9) {
                
                listoParaEnviar=false;
            }
            var input=dni;
            for (var i = input.length - 1; i >= 0; i--) {
                var letra_ilegal=false;
                var caracter_ilegal=false;
                for (var a = carateres_especiales.length - 1; a >= 0; a--) {
                    if (input[i]==carateres_especiales[a] || input[i]==arroba) {
                        listoParaEnviar=false;
                        caracter_ilegal=true;
                        break;                        
                    }                    

                }
                for (var o = caracteres_letras.length - 1; o >= 0; o--) {
                    if (input[i]==caracteres_letras[o]){
                        listoParaEnviar=false;
                        numero_ilegal=true;
                        break;
                    }
                }

                if (letra_ilegal) {
                    alert("El campo DNI tiene LETRAS ilegales, asegurate de usar solo numeros ");
                    break;
                }
                if (caracter_ilegal) {
                    alert("El campo DNI tiene caracteres ESPECIALES ilegales, asegurate de usar solo numeros ");
                    break;
                }

            }


            //Validar cuil
            if (cuil.length==0||cuil.length>12) {

                listoParaEnviar=false;
            }
             var input=cuil;
            for (var i = input.length - 1; i >= 0; i--) {
                var letra_ilegal=false;
                var caracter_ilegal=false;
                for (var a = carateres_especiales.length - 1; a >= 0; a--) {
                    if (input[i]==carateres_especiales[a] || input[i]==arroba) {
                        listoParaEnviar=false;
                        caracter_ilegal=true;
                        break;                        
                    }                    

                }
                for (var o = caracteres_letras.length - 1; o >= 0; o--) {
                    if (input[i]==caracteres_letras[o]){
                        listoParaEnviar=false;
                        numero_ilegal=true;
                        break;
                    }
                }

                if (letra_ilegal) {
                    alert("El campo CUIL tiene LETRAS ilegales, asegurate de usar solo numeros ");
                    break;
                }
                if (caracter_ilegal) {
                    alert("El campo CUIL tiene caracteres ESPECIALES ilegales, asegurate de usar solo numeros ");
                    break;
                }

            }
            //Validar fnacimiento
            if (fnacimiento.length==0) {
                listoParaEnviar=false;
            }

            //Validar email
            if (email.length==0||email.length>45) {

                listoParaEnviar=false;
            }

            var input=email;
            var con_arroba=false;

            for (var i = input.length - 1; i >= 0; i--) {

                if (input[i]==arroba[0]) {

                        con_arroba=true;
                        break;
                        
                }                   
            }

            if (!con_arroba) {
                alert("El campo EMAIIL no tiene formato de email");
                listoParaEnviar=false;
            }

            //Validar telefonos
            if(telefonofijo.length>11 || telefonoalter.length>11 || celular.length>11){
                listoParaEnviar=false;
            }
             var input=celular;
            for (var i = input.length - 1; i >= 0; i--) {
                var letra_ilegal=false;
                var caracter_ilegal=false;
                for (var a = carateres_especiales.length - 1; a >= 0; a--) {
                    if (input[i]==carateres_especiales[a] || input[i]==arroba) {
                        listoParaEnviar=false;
                        caracter_ilegal=true;
                        break;                        
                    }                    

                }
                for (var o = caracteres_letras.length - 1; o >= 0; o--) {
                    if (input[i]==caracteres_letras[o]){
                        listoParaEnviar=false;
                        numero_ilegal=true;
                        break;
                    }
                }

                if (letra_ilegal) {
                    alert("El campo CELULAR tiene LETRAS ilegales, asegurate de usar solo numeros ");
                    break;
                }
                if (caracter_ilegal) {
                    alert("El campo CELULAR tiene caracteres ESPECIALES ilegales, asegurate de usar solo numeros ");
                    break;
                }

            }
 
             var input=telefonoalter;
            for (var i = input.length - 1; i >= 0; i--) {
                var letra_ilegal=false;
                var caracter_ilegal=false;
                for (var a = carateres_especiales.length - 1; a >= 0; a--) {
                    if (input[i]==carateres_especiales[a] || input[i]==arroba) {
                        listoParaEnviar=false;
                        caracter_ilegal=true;
                        break;                        
                    }                    

                }
                for (var o = caracteres_letras.length - 1; o >= 0; o--) {
                    if (input[i]==caracteres_letras[o]){
                        listoParaEnviar=false;
                        numero_ilegal=true;
                        break;
                    }
                }

                if (letra_ilegal) {
                    alert("El campo TELEFONO ALTERNATIVO tiene LETRAS ilegales, asegurate de usar solo numeros ");
                    break;
                }
                if (caracter_ilegal) {
                    alert("El campo TELEFONO ALTERNATIVO tiene caracteres ESPECIALES ilegales, asegurate de usar solo numeros ");
                    break;
                }

            }
 
             var input=telefonofijo;
            for (var i = input.length - 1; i >= 0; i--) {
                var letra_ilegal=false;
                var caracter_ilegal=false;
                for (var a = carateres_especiales.length - 1; a >= 0; a--) {
                    if (input[i]==carateres_especiales[a] || input[i]==arroba) {
                        listoParaEnviar=false;
                        caracter_ilegal=true;
                        break;                        
                    }                    

                }
                for (var o = caracteres_letras.length - 1; o >= 0; o--) {
                    if (input[i]==caracteres_letras[o]){
                        listoParaEnviar=false;
                        numero_ilegal=true;
                        break;
                    }
                }

                if (letra_ilegal) {
                    alert("El campo TELEFONO FIJO tiene LETRAS ilegales, asegurate de usar solo numeros ");
                    break;
                }
                if (caracter_ilegal) {
                    alert("El campo TELEFONO FIJO tiene caracteres ESPECIALES ilegales, asegurate de usar solo numeros ");
                    break;
                }

            }
 
            
            
            
            //Validar direccion
            if (direccion.length==0||direccion.length>45) {
                listoParaEnviar=false;
            }
            var input=direccion;
            for (var i = input.length - 1; i >= 0; i--) {
                var numero_ilegal=false;
                var caracter_ilegal=false;
                for (var a = carateres_especiales.length - 1; a >= 0; a--) {
                    if (input[i]==carateres_especiales[a] || input[i]==arroba) {
                        listoParaEnviar=false;
                        caracter_ilegal=true;
                        break;                        
                    }                    

                }
                for (var o = caracteres_numeros.length - 1; o >= 0; o--) {
                    if (input[i]==caracteres_numeros[o]){
                        listoParaEnviar=false;
                        numero_ilegal=true;
                        break;
                    }
                }

                if (numero_ilegal) {
                    alert("El campo DIRECCION tiene caracteres NUMERICOS ilegales, asegurate de usar solo letras ");
                    break;
                }
                if (caracter_ilegal) {
                    alert("El campo DIRECCION tiene caracteres ESPECIALES ilegales, asegurate de usar solo letras ");
                    break;
                }

            }


            //Validar numero
            if (numero.length==0 || numero.length>11) {
                listoParaEnviar=false;
            }
             var input=numero;
            for (var i = input.length - 1; i >= 0; i--) {
                var letra_ilegal=false;
                var caracter_ilegal=false;
                for (var a = carateres_especiales.length - 1; a >= 0; a--) {
                    if (input[i]==carateres_especiales[a] || input[i]==arroba) {
                        listoParaEnviar=false;
                        caracter_ilegal=true;
                        break;                        
                    }                    

                }
                for (var o = caracteres_letras.length - 1; o >= 0; o--) {
                    if (input[i]==caracteres_letras[o]){
                        listoParaEnviar=false;
                        numero_ilegal=true;
                        break;
                    }
                }

                if (letra_ilegal) {
                    alert("El campo NUMERO tiene LETRAS ilegales, asegurate de usar solo numeros ");
                    break;
                }
                if (caracter_ilegal) {
                    alert("El campo NUMERO tiene caracteres ESPECIALES ilegales, asegurate de usar solo numeros ");
                    break;
                }

            }
 
            //valida piso y dpto
            if(piso.length>11 || dpto.length>45){
                listoParaEnviar=false;
            }
        
             var input=piso;
            for (var i = input.length - 1; i >= 0; i--) {
                var letra_ilegal=false;
                var caracter_ilegal=false;
                for (var a = carateres_especiales.length - 1; a >= 0; a--) {
                    if (input[i]==carateres_especiales[a] || input[i]==arroba) {
                        listoParaEnviar=false;
                        caracter_ilegal=true;
                        break;                        
                    }                    

                }
                for (var o = caracteres_letras.length - 1; o >= 0; o--) {
                    if (input[i]==caracteres_letras[o]){
                        listoParaEnviar=false;
                        numero_ilegal=true;
                        break;
                    }
                }

                if (letra_ilegal) {
                    alert("El campo PISO tiene LETRAS ilegales, asegurate de usar solo numeros ");
                    break;
                }
                if (caracter_ilegal) {
                    alert("El campo PISO tiene caracteres ESPECIALES ilegales, asegurate de usar solo numeros ");
                    break;
                }

            }
            var input=dpto;
            for (var i = input.length - 1; i >= 0; i--) {
                var numero_ilegal=false;
                var caracter_ilegal=false;
                for (var a = carateres_especiales.length - 1; a >= 0; a--) {
                    if (input[i]==carateres_especiales[a] || input[i]==arroba) {
                        listoParaEnviar=false;
                        caracter_ilegal=true;
                        break;                        
                    }                    

                }
               

                if (caracter_ilegal) {
                    alert("El campo DPTO tiene caracteres ESPECIALES ilegales, asegurate de usar letras y numeros ");
                    break;
                }

            }

            //Validar departamento
            if (departamento==0 && provincia==0) {
                listoParaEnviar=false;
            }


            if(listoParaEnviar){
                app.verificarExistenciaPersona();
            }

        }; 

        /*app.verificarExistenciaDomicilio = function() {
            var url = "controlador/Ruteador.php?accion=ExisteDomi&id=-1";

            var direccion=document.getElementById('direccion').value;
            var numero=document.getElementById('numero').value;
            var depto=document.getElementById('depto').value;
            var piso=document.getElementById('piso').value;

             e = document.getElementById("departamento");
            var departamento = e.options[e.selectedIndex].value;


            $.ajax({
                url: url,
                method: 'POST',
                dataType: 'json',
                data: {Direccion: direccion,
                    Numero: numero,
                    Piso: piso,
                    Dpto:depto,
                    Departamento: departamento},
                success: function (datosRecibidos) {
                        if(datosRecibidos=="0"){
                        app.verificarExistenciaPersona();
                        } else{
                            alert("El domicilio ingresado ya existe en su departamento");
                        }
                },
                error: function (datosRecibidos) {
                    alert("Error de busqueda");

                }
            });
        };*/

        app.verificarExistenciaPersona = function(){
             var url = "controlador/Ruteador.php?accion=ExistePersona&id=-1";

            var dni=document.getElementById('dni').value;

            $.ajax({
                url: url,
                method: 'POST',
                dataType: 'json',
                data: {DNI:dni},
                success: function (datosRecibidos) {

                    if(datosRecibidos=="0"){
                        app.guardarRegistro();
                        } else{
                            alert("Usted ya esta dentro registrado en el sistema");
                        }

                    
                },
                error: function (datosRecibidos) {
                    alert("Error de busqueda");

                }
            });
        };

        // llamada ajax para guardar alumno
        app.guardarRegistro = function () {  //funcion para guardar un alumno
            var url = "controlador/Ruteador.php?accion=agregar&id=-1";
            // variable que toma todos los datos del formulario
            var estCivil = $('input:radio[name=estado_civil]:checked').val();
            var genero = $('input:radio[name=genero]:checked').val();
            e = document.getElementById("departamento");
            var departamento = e.options[e.selectedIndex].value;
            var nombre=document.getElementById('nombre').value;
            var apellido=document.getElementById('apellido').value;
            var apellidomaterno=document.getElementById('apellidomaterno').value;
            var dni=document.getElementById('dni').value;
            var cuil=document.getElementById('cuil').value;
            var fnacimiento=document.getElementById('fecha_nacimiento').value;
            var email=document.getElementById('correo').value;
            var telefono=document.getElementById('telefono').value;
            var celular=document.getElementById('celular').value;
            var telalter=document.getElementById('telefonoa').value;
            var direccion=document.getElementById('direccion').value;
            var numero=document.getElementById('numero').value;
            var piso=document.getElementById('piso').value;
            var depto=document.getElementById('depto').value;





            var enviarD = {
                    Nombre: nombre,
                    Apellido: apellido,
                    ApellidoMaterno: apellidomaterno,
                    DNI: dni,
                    Cuil: cuil,
                    FechaNacimiento: fnacimiento,
                    Email: email,
                    EstadoCivil: estCivil,
                    Genero: genero,
                    Telefono: telefono,
                    Celular: celular,
                    TelefonoAlternativo: telalter,
                    Direccion: direccion,
                    Numero: numero,
                    Piso: piso,
                    Dpto:depto,
                    Departamento: departamento
                
            };

            $.ajax({
                url: url,
                method: 'POST',
                dataType: 'json',
                data: enviarD,
                success: function (datosRecibidos) {
                    // esconder el modal
                    alert("Formulario creado correctamente");
                    app.limpiarFormulario(); //Limpiar Formulario
                },
                error : function(XMLHttpRequest, textStatus, errorThrown) {
                        alert(XMLHttpRequest);
                        alert(textStatus);
                        alert(errorThrown);
                        alert(XMLHttpRequest.responseText);
               }
            });
        };

        
        app.actualizarProvincia = function (data) {

            var html = '<option value="0" selected>Seleccione... </option>';

            var x = 0;


            for (var i = data.length - 1; i >= 0; i--) {

            html += "<option value='" + data[x].OID_provincia + "'>" + data[x].nombre + "</option>" + "\n";
            x+=1;
            }
            

            document.getElementById("provincia").innerHTML = html;
        };

        app.actualizarDepartamento = function (data) {

            var html = '';

            var x = 0;

            for (var i = data.length - 1; i >= 0; i--) {

            html += "<option value='" + data[x].OID_departamento + "'>" + data[x].nombre + "</option>" + "\n";
            x+=1;
            }
            

            document.getElementById("departamento").innerHTML = html;
        }

        app.init();

    })(espacioNombre);


});