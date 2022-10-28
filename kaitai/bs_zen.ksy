meta:
  id: bs_zen
  file-extension: zen
  endian: le
doc: |
  File format definition for binary safe ZEN world archive format for the ZenGin.
  
seq:
  - id: header
    type: header
    doc: |
      Archive header
  - id: world
    type: world
  - id: o_c_world
    type: o_c_world
    #9241
    #repeat: eos #until
    #repeat-until: _.pos == _parent.bsp #9243

    

types:
  header:
    seq:
      - id: magic
        type: string
      - id: version
        type: string
      - id: archiver_class
        type: string
      - id: archiver_type
        type: string
      - id: save_game
        type: string
      - id: date
        type: string
      - id: user
        type: string
      - id: end
        type: string

  property:
    seq:
      - id: type
        type: u1
        #enum: prop_type
      - id: prop_body
        type:
          switch-on: type
          cases:
            0x1: type_string
            0x2: type_integer
            0x3: type_float
            0x4: type_byte
            0x5: type_word
            0x6: type_bool
            0x7: type_vec3
            0x8: type_color
            0x9: type_raw
            0x10: type_raw_float
            0x11: type_enum
            0x12: type_hash
  o_c_world:
    seq:
      - id: object_id
        type: property
      - id: mab
        type: mab
      - id: properties
        type: property
        repeat: until
        repeat-until: _parent.world.hash_table_offset - _io.pos == 0
        
  world:
    seq:
      - id: bs_version
        type: u4
      - id: object_count
        type: u4
      - id: hash_table_offset
        type: u4
    instances:
      ht:
        pos: hash_table_offset
        type: hash_table_header
        
  mab:
    seq:
      - id: mesh_and_bsp
        type: property
      - id: bsp
        type: bsp

  type_string:
    seq:
      - id: size
        type: u2
      - id: object_id
        type: str
        size: size
        encoding: ASCII

  type_integer:
    seq:
      - id: value
        type: u4
  
  type_float:
    seq:
      - id: value
        type: f4
  
  type_byte:
    seq:
      - id: value
        size: 1
  
  type_word:
    seq:
      - id: value
        size: 2
  
  type_bool:
    seq:
      - id: value
        type: u4

  type_vec3:
    seq:
      - id: x
        type: f4
      - id: y
        type: f4
      - id: z
        type: f4
  
  type_color:
    seq:
      - id: r
        type: u1
      - id: g
        type: u1
      - id: b
        type: u1
      - id: a
        type: u1
  
  type_raw:
    seq:
      - id: size
        type: u2
      - id: raw_data
        size: size
  
  type_raw_float:
    seq:
      - id: size
        type: u2
      - id: raw_float_data
        size: size
  
  type_enum:
    seq:
      - id: value
        type: u4
  
  type_hash:
    seq:
      - id: value
        type: u4

  hash_table_header:
    seq:
      - id: chunk_count
        type: u4
      - id: chunks
        type: hash_table_chunk
        repeat: expr
        repeat-expr: chunk_count

  hash_table_chunk:
    seq:
      - id: name_length
        type: u2
      - id: linear_value
        type: u2
      - id: hash_value
        type: u4
      - id: str
        type: str
        size: name_length
        encoding: ASCII

  string:
    seq:
      - id: str
        type: str
        terminator: 0xa
        encoding: ASCII


  bsp:
    seq:
      - id: bsp_magic
        type: u4
      - id: size
        type: u4
      - id: version
        type: u4
      - id: data
        size: size + 1

    
    

        